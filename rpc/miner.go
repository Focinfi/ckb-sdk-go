package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types/errtypes"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func (client *Client) GetBlockTemplate(ctx context.Context, bytesLimit *types.HexUint64, proposalsLimit *types.HexUint64, maxVersion *types.HexUint64) (*ckbtypes.BlockTemplate, error) {
	params := []interface{}{optionHex(bytesLimit), optionHex(proposalsLimit), optionHex(maxVersion)}
	result, err := RawHTTPPost(ctx, client.URL, "get_epoch_by_number", params)
	if err != nil {
		return nil, err
	}

	var template ckbtypes.BlockTemplate
	if err := json.Unmarshal(result, &template); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &template, nil
}

func (client *Client) SubmitBlock(ctx context.Context, workID string, block ckbtypes.Block) (string, error) {
	params := []interface{}{workID, block}
	result, err := RawHTTPPost(ctx, client.URL, "submit_block", params)
	if err != nil {
		return "", err
	}

	var blockHash string
	if err := json.Unmarshal(result, &blockHash); err != nil {
		return "", errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return blockHash, nil
}
