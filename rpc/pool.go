package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client *Client) SendTransaction(ctx context.Context, transaction ckbtypes.RawTransaction) (string, error) {
	params := []interface{}{transaction}
	result, err := RawHTTPPost(ctx, client.URL, "send_transaction", params)
	if err != nil {
		return "", err
	}

	var txHash string
	if err := json.Unmarshal(result, &txHash); err != nil {
		return "", errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return txHash, nil
}

func (client *Client) TxPoolInfo(ctx context.Context) (*ckbtypes.PoolInfo, error) {
	result, err := RawHTTPPost(ctx, client.URL, "tx_pool_info", nil)
	if err != nil {
		return nil, err
	}

	var poolInfo ckbtypes.PoolInfo
	if err := json.Unmarshal(result, &poolInfo); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &poolInfo, nil
}
