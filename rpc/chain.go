package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client *Client) GetTipBlockNumber(ctx context.Context) (uint64, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_tip_block_number", nil)
	if err != nil {
		return 0, err
	}

	var hexUint64 types.HexUint64
	if err := json.Unmarshal(result, &hexUint64); err != nil {
		return 0, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return hexUint64.Uint64(), nil
}

func (client *Client) GetTipHeader(ctx context.Context) (*ckbtypes.Header, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_tip_header", nil)
	if err != nil {
		return nil, err
	}

	var header ckbtypes.Header
	if err := json.Unmarshal(result, &header); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &header, nil
}

func (client *Client) GetCurrentEpoch(ctx context.Context) (*ckbtypes.Epoch, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_current_epoch", nil)
	if err != nil {
		return nil, err
	}

	var epoch ckbtypes.Epoch
	if err := json.Unmarshal(result, &epoch); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &epoch, nil
}

func (client *Client) GetEpochByNumber(ctx context.Context, num types.HexUint64) (*ckbtypes.Epoch, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_epoch_by_number", []string{num.Hex()})
	if err != nil {
		return nil, err
	}

	var epoch ckbtypes.Epoch
	if err := json.Unmarshal(result, &epoch); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &epoch, nil
}

func (client *Client) GetBlockHash(ctx context.Context, num types.HexUint64) (string, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_block_hash", []string{num.Hex()})
	if err != nil {
		return "", err
	}

	var hash string
	if err := json.Unmarshal(result, &hash); err != nil {
		return "", errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return hash, nil
}

func (client *Client) GetBlock(ctx context.Context, blockHash string) (*ckbtypes.Block, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_block", []string{blockHash})
	if err != nil {
		return nil, err
	}

	var block ckbtypes.Block
	if err := json.Unmarshal(result, &block); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &block, nil
}

func (client *Client) GetBlockByNumber(ctx context.Context, num types.HexUint64) (*ckbtypes.Block, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_block_by_number", []string{num.Hex()})
	if err != nil {
		return nil, err
	}

	var block ckbtypes.Block
	if err := json.Unmarshal(result, &block); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &block, nil
}

func (client *Client) GetHeader(ctx context.Context, blockHash string) (*ckbtypes.Header, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_header", []string{blockHash})
	if err != nil {
		return nil, err
	}

	var header ckbtypes.Header
	if err := json.Unmarshal(result, &header); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &header, nil
}

func (client *Client) GetHeaderByNumber(ctx context.Context, num types.HexUint64) (*ckbtypes.Header, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_header_by_number", []string{num.Hex()})
	if err != nil {
		return nil, err
	}

	var header ckbtypes.Header
	if err := json.Unmarshal(result, &header); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &header, nil
}

func (client *Client) GetCellsByLockHash(ctx context.Context, lockHash string, from, to types.HexUint64) ([]ckbtypes.CellOutputWithOutPoint, error) {
	params := []string{lockHash, from.Hex(), to.Hex()}
	result, err := RawHTTPPost(ctx, client.URL, "get_cells_by_lock_hash", params)
	if err != nil {
		return nil, err
	}

	var cells []ckbtypes.CellOutputWithOutPoint
	if err := json.Unmarshal(result, &cells); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return cells, nil
}

func (client *Client) GetLiveCell(ctx context.Context, outpoint ckbtypes.OutPoint, fetchData bool) (*ckbtypes.CellWithStatus, error) {
	params := []interface{}{outpoint, fetchData}
	result, err := RawHTTPPost(ctx, client.URL, "get_live_cell", params)
	if err != nil {
		return nil, err
	}

	var cellInfo ckbtypes.CellWithStatus
	if err := json.Unmarshal(result, &cellInfo); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &cellInfo, nil
}

func (client *Client) GetTransaction(ctx context.Context, txHash string) (*ckbtypes.TransactionInfo, error) {
	params := []interface{}{txHash}
	result, err := RawHTTPPost(ctx, client.URL, "get_transaction", params)
	if err != nil {
		return nil, err
	}

	var transactionInfo ckbtypes.TransactionInfo
	if err := json.Unmarshal(result, &transactionInfo); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &transactionInfo, nil
}

func (client *Client) GetCellBaseOutputCapacityDetails(ctx context.Context, blockHash string) (*ckbtypes.CellBase, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_cellbase_output_capacity_details", []string{blockHash})
	if err != nil {
		return nil, err
	}

	var cellBase ckbtypes.CellBase
	if err := json.Unmarshal(result, &cellBase); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &cellBase, nil
}
