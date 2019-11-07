package wallet

import (
	"context"
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/cellcollector"
	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/utils"
)

type Wallet struct {
	Client          *rpc.Client
	Key             *key.Key
	SkipDataAndType bool
	lockHashHex     *types.HexStr
	lock            *ckbtypes.Script
}

func NewWallet(client *rpc.Client, key *key.Key, skipDataAndType bool) (*Wallet, error) {
	lockHashHex, err := utils.LockScriptHash(key.Address.PubKey.PubKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Client:          client,
		Key:             key,
		SkipDataAndType: skipDataAndType,
		lockHashHex:     lockHashHex,
	}, nil
}

func NewWalletByPrivKey(client *rpc.Client, privKey string, skipDataAndType bool, mode types.Mode) (*Wallet, error) {
	k, err := key.NewFromPrivKeyHex(privKey, mode)
	if err != nil {
		return nil, err
	}
	return NewWallet(client, k, skipDataAndType)
}

func (wallet *Wallet) Balance(ctx context.Context) (uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	_, totalCap, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), 0)
	if err != nil {
		return 0, err
	}
	return totalCap, nil
}

func (wallet *Wallet) GetUnspentCells(ctx context.Context, needCap uint64) ([]ckbtypes.Cell, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	cells, _, err := collector.GetUnspentCells(ctx, wallet.lockHashHex.Hex(), needCap)
	if err != nil {
		return nil, err
	}
	return cells, nil
}

func (wallet *Wallet) GenerateTx(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64, useDepGroup bool) (*ckbtypes.Transaction, error) {
	arg, err := address.ParseShortPayloadAddressArg(targetAddr, wallet.Key.Address.Mode)
	if err != nil {
		return nil, err
	}
	dataHex := types.NewHexStr(data)
	output := ckbtypes.Output{
		Capacity: types.HexUint64(capacity).Hex(),
		Lock: ckbtypes.Script{
			Args:     arg,
			CodeHash: types.BlockAssemblerCodeHash,
			HashType: ckbtypes.HashTypeType,
		},
	}
	outputByteSize, err := output.ByteSize()
	if err != nil {
		return nil, err
	}
	changeOutput := ckbtypes.Output{Lock: *wallet.Lock()}
	minChangeByteSize, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}
	minCap := (outputByteSize + uint64(dataHex.Len())) * types.OneCKBShannon
	minChangeCap := (minChangeByteSize + uint64(dataHex.Len())) * types.OneCKBShannon
	inputs, inputCap, err := wallet.GatherInputs(ctx, capacity, minCap, minChangeCap, fee)
	if err != nil {
		return nil, err
	}
	outputs := []ckbtypes.Output{output}
	outputsData := []string{dataHex.Hex()}
	if changeCap := inputCap - (capacity + fee); changeCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}
	tx := &ckbtypes.Transaction{
		Version:     types.HexUint64(0).Hex(),
		CellDeps:    []ckbtypes.CellDep{},
		HeaderDeps:  []string{},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   utils.EmptyWitnessesByLen(len(inputs)),
	}

	sysCells, err := utils.LoadSystemCells(*wallet.Client)
	if err != nil {
		return nil, err
	}
	if useDepGroup {
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeDepGroup, OutPoint: *sysCells.Secp256k1GroupOutPoint})
	} else {
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeCode, OutPoint: *sysCells.Secp256k1CodeOutPoint})
		tx.CellDeps = append(tx.CellDeps, ckbtypes.CellDep{DepType: ckbtypes.DepTypeCode, OutPoint: *sysCells.Secp256k1CodeOutPoint})
	}
	if err := utils.SignTransaction(*wallet.Key, tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (wallet *Wallet) SendCapacity(ctx context.Context, targetAddr string, capacity uint64, data []byte, fee uint64) (string, error) {
	tx, err := wallet.GenerateTx(ctx, targetAddr, capacity, data, fee, true)
	if err != nil {
		return "", nil
	}
	return wallet.SendTransaction(ctx, *tx)
}

func (wallet *Wallet) DepositToDAO(ctx context.Context, capacity, fee uint64) (*ckbtypes.OutPoint, error) {
	sysCells, err := utils.LoadSystemCells(*wallet.Client)
	if err != nil {
		return nil, err
	}
	output := ckbtypes.Output{
		Capacity: types.HexUint64(capacity).Hex(),
		Lock:     *wallet.Lock(),
		Type: &ckbtypes.Script{
			Args:     types.HexStrPrefix,
			CodeHash: sysCells.DaoTypeHash.Hex(),
			HashType: ckbtypes.HashTypeType,
		},
	}
	outputByteSize, err := output.ByteSize()
	if err != nil {
		return nil, err
	}
	changeOutput := ckbtypes.Output{
		Lock: *wallet.Lock(),
	}
	changeOutputByteSize, err := changeOutput.ByteSize()
	if err != nil {
		return nil, err
	}
	minCap := outputByteSize * types.OneCKBShannon
	minChangeCap := changeOutputByteSize * types.OneCKBShannon
	inputs, inputsCap, err := wallet.GatherInputs(ctx, capacity, minCap, minChangeCap, fee)
	outputs := []ckbtypes.Output{output}
	outputsData := []string{types.HexStrPrefix}
	changeOutputCap := inputsCap - (capacity + fee)
	if changeOutputCap > 0 {
		changeOutput.Capacity = types.HexUint64(changeOutputCap).Hex()
		outputs = append(outputs, changeOutput)
		outputsData = append(outputsData, types.HexStrPrefix)
	}
	tx := ckbtypes.Transaction{
		Version: types.HexUint64(0).Hex(),
		CellDeps: []ckbtypes.CellDep{
			{OutPoint: *sysCells.Secp256k1GroupOutPoint, DepType: ckbtypes.DepTypeDepGroup},
			{OutPoint: *sysCells.DaoOutPoint, DepType: ckbtypes.DepTypeCode},
		},
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Witnesses:   utils.EmptyWitnessesByLen(len(inputs)),
	}
	if err := utils.SignTransaction(*wallet.Key, &tx); err != nil {
		return nil, err
	}
	txHash, err := wallet.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return &ckbtypes.OutPoint{
		TxHash: txHash,
		Index:  types.HexUint64(0).Hex(),
	}, nil
}

func (wallet *Wallet) GenerateWithdrawFromDAOTransaction(ctx context.Context, point ckbtypes.OutPoint, fee uint64) (*ckbtypes.Transaction, error) {
	return nil, nil
}

func (wallet *Wallet) GetTransaction(ctx context.Context, hash string) (*ckbtypes.TransactionInfo, error) {
	return wallet.Client.GetTransaction(ctx, hash)
}

func (wallet *Wallet) BlockAssemblerConfig() string {
	return fmt.Sprintf(
		`[block_assembler]
	code_hash = %s
	args = %s`,
		wallet.Lock().CodeHash,
		wallet.Lock().Args)
}

func (wallet *Wallet) SendTransaction(ctx context.Context, transaction ckbtypes.Transaction) (string, error) {
	return wallet.Client.SendTransaction(ctx, transaction.ToRaw())
}

func (wallet *Wallet) GatherInputs(ctx context.Context, capacity, minCap, minChangeCap, fee uint64) ([]ckbtypes.Input, uint64, error) {
	collector := cellcollector.NewCellCollector(wallet.Client, wallet.SkipDataAndType)
	return collector.GatherInputs(ctx, []string{wallet.lockHashHex.Hex()}, capacity, minChangeCap, minChangeCap, fee)
}

func (wallet *Wallet) Lock() *ckbtypes.Script {
	if wallet.lock == nil {
		wallet.lock = &ckbtypes.Script{
			Args:     wallet.Key.Address.PubKey.Blake160.Hex(),
			CodeHash: types.BlockAssemblerCodeHash,
			HashType: ckbtypes.HashTypeType,
		}
	}
	return wallet.lock
}

func (wallet *Wallet) CodeHash(ctx context.Context) string {
	return wallet.Lock().CodeHash
}
