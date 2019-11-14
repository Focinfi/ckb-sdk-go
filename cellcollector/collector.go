package cellcollector

import (
	"context"
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

// CellCollector represents a collector to collect the cells via rpc
type CellCollector struct {
	Cli             *rpc.Client
	SkipDataAndType bool
}

func NewCellCollector(cli *rpc.Client, skipDataAndType bool) *CellCollector {
	return &CellCollector{
		Cli:             cli,
		SkipDataAndType: skipDataAndType,
	}
}

// GetUnspentCells gets unspent cells.
// Get all unspent cells when needCap <= 0
// The given needCap in shanon
func (collector *CellCollector) GetUnspentCells(ctx context.Context, lockHash string, needCap uint64) ([]ckbtypes.CellOutputWithOutPoint, uint64, error) {
	tipBlockNum, err := collector.Cli.GetTipBlockNumber(ctx)
	if err != nil {
		return nil, 0, err
	}

	var (
		unspentCells []ckbtypes.CellOutputWithOutPoint
		currentFrom  uint64
		totalCap     uint64
	)
	for currentFrom <= tipBlockNum {
		currentTo := currentFrom + 100
		if currentTo > tipBlockNum {
			currentTo = tipBlockNum
		}
		cells, err := collector.Cli.GetCellsByLockHash(ctx, lockHash, types.HexUint64(currentFrom), types.HexUint64(currentTo))
		if err != nil {
			return nil, 0, err
		}

		for _, cell := range cells {
			if collector.SkipDataAndType {
				liveCell, err := collector.Cli.GetLiveCell(ctx, cell.OutPoint, true)
				if err != nil {
					return nil, 0, err
				}
				data := liveCell.Cell.Data.Content

				isNilDataAndType := (data == "" || data == types.HexStrPrefix) && liveCell.Cell.Output.Type == nil
				if !isNilDataAndType {
					continue
				}
			}

			unspentCells = append(unspentCells, cell)
			hexNum, err := types.ParseHexUint64(cell.Capacity)
			if err != nil {
				return nil, 0, err
			}
			totalCap += hexNum.Uint64()

			if needCap > 0 && totalCap >= needCap {
				return unspentCells, totalCap, nil
			}
		}

		currentFrom = currentTo + 1
	}

	return unspentCells, totalCap, nil
}

// GetUnspentCellsByLockHashes gets unspent cells by lock hashes
func (collector *CellCollector) GetUnspentCellsByLockHashes(ctx context.Context, lockHashes []string, needCap uint64) ([]ckbtypes.CellOutputWithOutPoint, uint64, error) {
	var (
		totalCap     uint64
		unspentCells []ckbtypes.CellOutputWithOutPoint
	)
	for _, lockHash := range lockHashes {
		cells, c, err := collector.GetUnspentCells(ctx, lockHash, needCap)
		if err != nil {
			return nil, 0, err
		}
		unspentCells = append(unspentCells, cells...)
		totalCap += c
		if totalCap >= needCap {
			return unspentCells, totalCap, nil
		}
	}

	return unspentCells, totalCap, nil
}

// GatherInputs gathers the inputs.
// The given needCap <= minCap.
// The given needCap + fee >= all unspent cell capacity
func (collector *CellCollector) GatherInputs(ctx context.Context, lockHashes []string, needCap, minCap, minChangeCap, fee uint64) ([]ckbtypes.Input, uint64, error) {
	if needCap < minCap {
		return nil, 0, errtypes.WrapErr(errtypes.RPCErrGetInputCapacityTooSmall, fmt.Errorf("need capactiy is less than %d", minCap))
	}

	var (
		totalCap = needCap + fee
		inputCap uint64
		inputs   []ckbtypes.Input
	)

	cells, _, err := collector.GetUnspentCellsByLockHashes(ctx, lockHashes, needCap)
	if err != nil {
		return nil, 0, err
	}
	for _, cell := range cells {
		input := ckbtypes.Input{
			PreviousOutput: cell.OutPoint,
			Since:          types.HexUint64(0).Hex(),
		}
		hexNum, err := types.ParseHexUint64(cell.Capacity)
		if err != nil {
			return nil, 0, err
		}
		inputs = append(inputs, input)
		inputCap += hexNum.Uint64()
		if inputCap >= totalCap+minChangeCap || inputCap == totalCap {
			return inputs, inputCap, nil
		}
	}

	if inputCap < totalCap {
		return nil, 0, errtypes.WrapErr(errtypes.RPCErrGetInputCapacityNotEnough, nil)
	}
	return inputs, inputCap, nil
}
