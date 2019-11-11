package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client *Client) GetBlockchainInfo(ctx context.Context) (*ckbtypes.BlockChainInfo, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_blockchain_info", nil)
	if err != nil {
		return nil, err
	}

	var blockchainInfo ckbtypes.BlockChainInfo
	if err := json.Unmarshal(result, &blockchainInfo); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &blockchainInfo, nil
}

func (client *Client) GetPeersState(ctx context.Context) ([]ckbtypes.PeerState, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_peers_state", nil)
	if err != nil {
		return nil, err
	}

	var states []ckbtypes.PeerState
	if err := json.Unmarshal(result, &states); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return states, nil
}
