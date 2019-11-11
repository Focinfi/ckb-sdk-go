package rpc

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func (client *Client) LockNodeInfo(ctx context.Context) (*ckbtypes.NodeInfo, error) {
	result, err := RawHTTPPost(ctx, client.URL, "local_node_info", nil)
	if err != nil {
		return nil, err
	}

	var nodeInfo ckbtypes.NodeInfo
	if err := json.Unmarshal(result, &nodeInfo); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return &nodeInfo, nil
}

func (client *Client) GetPeers(ctx context.Context) ([]ckbtypes.NodeInfo, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_peers", nil)
	if err != nil {
		return nil, err
	}

	var nodeInfos []ckbtypes.NodeInfo
	if err := json.Unmarshal(result, &nodeInfos); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return nodeInfos, nil
}

func (client *Client) GetBannedAddresses(ctx context.Context) ([]ckbtypes.BannedAddress, error) {
	result, err := RawHTTPPost(ctx, client.URL, "get_banned_addresses", nil)
	if err != nil {
		return nil, err
	}

	var addresses []ckbtypes.BannedAddress
	if err := json.Unmarshal(result, &addresses); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResponseDataBroken, err)
	}

	return addresses, nil
}

func (client *Client) SetBan(ctx context.Context, address, command string, banTime int64, absolute bool, reason string) error {
	params := []interface{}{address, command, types.HexUint64(banTime), absolute, reason}
	_, err := RawHTTPPost(ctx, client.URL, "set_ban", params)
	return err
}
