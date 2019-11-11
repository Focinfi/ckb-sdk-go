package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client Client) DryRunTransaction(ctx context.Context, transaction ckbtypes.Transaction) (uint64, error) {
	result, err := RawHTTPPost(ctx, client.URL, "dry_run_transaction", nil)
	if err != nil {
		return 0, err
	}

	var hexUint64 types.HexUint64
	if err := json.Unmarshal(result, &hexUint64); err != nil {
		return 0, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return hexUint64.Uint64(), nil
}

func (client *Client) CalculateDAOMaximumWithdraw(ctx context.Context, outpoint ckbtypes.OutPoint, withdrawBlockHash string) (uint64, error) {
	params := []interface{}{outpoint, withdrawBlockHash}
	result, err := RawHTTPPost(ctx, client.URL, "calculate_dao_maximum_withdraw", params)
	if err != nil {
		return 0, err
	}

	var hexUint64 types.HexUint64
	if err := json.Unmarshal(result, &hexUint64); err != nil {
		return 0, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return hexUint64.Uint64(), nil
}
