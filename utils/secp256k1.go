package utils

import (
	"context"
	"errors"

	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

var secp256k1OutPoint *ckbtypes.OutPoint
var secp256k1ScriptHash *types.HexStr

func GetSecp256k1OutPointAndScriptHash(client *rpc.Client) (*ckbtypes.OutPoint, *types.HexStr, error) {
	if secp256k1OutPoint != nil {
		return secp256k1OutPoint, nil, nil
	}

	block, err := client.GetBlockByNumber(context.Background(), 0)
	if err != nil {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, err)
	}
	if len(block.Transactions) < 2 {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("transaction length < 2"))
	}
	if len(block.Transactions[0].Outputs) < 2 {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("transactions[0].outputs length < 2"))
	}
	txHash := block.Transactions[1].Hash
	typeScript := block.Transactions[0].Outputs[1].Type
	if txHash == "" {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("transaction[1].hash is empty"))
	}
	if typeScript == nil {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("transaction[0].outputs[1].type is empty"))
	}
	secp256k1ScriptHash, err = ScriptHash(*typeScript)
	if err != nil {
		return nil, nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, err)
	}
	secp256k1OutPoint = &ckbtypes.OutPoint{
		TxHash: txHash,
		Index:  types.HexUint64(0).Hex(),
	}
	return secp256k1OutPoint, secp256k1ScriptHash, nil
}
