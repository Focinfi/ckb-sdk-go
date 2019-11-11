package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client *Client) IndexLockHash(ctx context.Context, lockHash string, indexFrom types.HexUint64) (*ckbtypes.BlockHashIndexState, error) {
	params := []interface{}{lockHash, indexFrom.Hex()}
	result, err := RawHTTPPost(ctx, client.URL, "index_lock_hash", params)
	if err != nil {
		return nil, err
	}

	var state ckbtypes.BlockHashIndexState
	if err := json.Unmarshal(result, &state); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &state, nil
}

func (client *Client) GetLockHashIndexStats(ctx context.Context) (*ckbtypes.BlockHashIndexState, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_lock_hash_index_states", nil)
	if err != nil {
		return nil, err
	}

	var state ckbtypes.BlockHashIndexState
	if err := json.Unmarshal(result, &state); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &state, nil
}

func (client *Client) GetLiveCellsByLockHash(ctx context.Context, lockHash string, page, per types.HexUint64, reverseOrder bool) ([]ckbtypes.LiveCell, error) {
	params := []interface{}{lockHash, page.Hex(), per.Hex(), reverseOrder}
	result, err := RawHTTPPost(ctx, client.URL, "get_live_cells_by_lock_hash", params)
	if err != nil {
		return nil, err
	}

	var cells []ckbtypes.LiveCell
	if err := json.Unmarshal(result, &cells); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return cells, nil
}

func (client *Client) GetTransactionByLockHash(ctx context.Context, lockHash string, page, per types.HexUint64, reverseOrder bool) ([]ckbtypes.CellTransaction, error) {
	params := []interface{}{lockHash, page.Hex(), per.Hex(), reverseOrder}
	result, err := RawHTTPPost(ctx, client.URL, "get_transactions_by_lock_hash", params)
	if err != nil {
		return nil, err
	}

	var transactions []ckbtypes.CellTransaction
	if err := json.Unmarshal(result, &transactions); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return transactions, nil
}

func (client *Client) DeindexLockHash(ctx context.Context, lockHash string) error {
	params := []interface{}{lockHash}
	_, err := RawHTTPPost(ctx, client.URL, "deindex_lock_hash", params)
	return err
}
