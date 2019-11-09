package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type SysCells struct {
	Secp256k1CodeHash      *types.HexStr
	Secp256k1CodeOutPoint  *ckbtypes.OutPoint
	Secp256k1DataOutPoint  *ckbtypes.OutPoint
	Secp256k1GroupOutPoint *ckbtypes.OutPoint

	DaoOutPoint *ckbtypes.OutPoint
	DaoCodeHash *types.HexStr
	DaoTypeHash *types.HexStr

	MultiSignSecpCellTypeHash  *types.HexStr
	MultiSignSecpGroupOutPoint *ckbtypes.OutPoint
}

var sysCells *SysCells

// LoadSystemCells loads genesis block and extract the secp256k1 and DAO data
// secp256k1
// 	systemCellTrans = blocks[0].transaction[0]
// 	secp256k1Code = systemCellTrans.output_data[1]
// 	secp256k1CodeOutPoint = { index: 1, tx_hash: systemCellTrans.hash }
// 	secp256k1DataOutPoint = { index: 3, tx_hash: systemCellTrans.hash }
//
// 	secp256k1GroupTransaction = blocks[0].transaction[1]
// 	secp256k1GroupOutPoint = { index: 0, tx_hash: secp256k1GroupTransaction.hash }
// DAO
//  daoOutPoint = { index: 2, tx_hash: systemCellTrans.hash }
//  daoCode = systemCellTrans.output_data[2]
//  daoTypeHash = systemCellTrans.outputs[2].type.hash
// Multi Sign
//  hash = systemCellTrans.outputs[4].type.compute_hash
//  out point = { tx_hash: secp256k1GroupTransaction.hash, index: 1 }
func LoadSystemCells(client rpc.Client) (*SysCells, error) {
	if sysCells != nil {
		return sysCells, nil
	}
	block, err := client.GetBlockByNumber(context.Background(), 0)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, err)
	}
	if len(block.Transactions) < 2 {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("genesis block trans length < 2"))
	}
	if len(block.Transactions[0].Outputs) < 5 {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("systemCellTrans.outputs length < 5"))
	}
	if len(block.Transactions[0].OutputsData) < 3 {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("systemCellTrans.outputs_data length < 3"))
	}
	// secp256k1
	sysCellTrans := block.Transactions[0]
	secp256k1Code := sysCellTrans.OutputsData[1]
	secp256k1CodeHex, err := types.ParseHexStr(secp256k1Code)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("decode secp256k1 code hex failed: %v", err))
	}
	secp256k1CodeHash, err := blake2b.Digest(secp256k1CodeHex.Bytes())
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("hash secp256k1 code failed: %v", err))
	}

	secp256k1CodeOutPoint := &ckbtypes.OutPoint{
		Index:  types.HexUint64(1).Hex(),
		TxHash: sysCellTrans.Hash,
	}
	secp256k1DataOutPoint := &ckbtypes.OutPoint{
		Index:  types.HexUint64(3).Hex(),
		TxHash: sysCellTrans.Hash,
	}
	secp256k1GroupTrans := block.Transactions[1]
	secp256k1GroupOutPoint := &ckbtypes.OutPoint{
		Index:  types.HexUint64(0).Hex(),
		TxHash: secp256k1GroupTrans.Hash,
	}
	// DAO
	daoOutPoint := &ckbtypes.OutPoint{
		Index:  types.HexUint64(2).Hex(),
		TxHash: sysCellTrans.Hash,
	}
	daoCode := sysCellTrans.OutputsData[2]
	daoCodeHex, err := types.ParseHexStr(daoCode)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("decode dao code hex failed: %v", err))
	}
	daoCodeHash, err := blake2b.Digest(daoCodeHex.Bytes())
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("hash dao code failed: %v", err))
	}
	daoTypeHash, err := ScriptHash(*sysCellTrans.Outputs[2].Type)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("hash dao type script failed: %v", err))
	}

	// Multi Sign
	if sysCellTrans.Outputs[4].Type == nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, errors.New("systemCellTrans.outputs[4].type is nil"))
	}
	multiSignTypeHash, err := ScriptHash(*sysCellTrans.Outputs[4].Type)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrGetGenesisBlockBroken, fmt.Errorf("hash multisign type script failed: %v", err))
	}
	multiSignSecpGroupOutPoint := ckbtypes.OutPoint{
		Index:  types.HexUint64(1).Hex(),
		TxHash: secp256k1GroupTrans.Hash,
	}

	sysCells = &SysCells{
		Secp256k1CodeHash:      types.NewHexStr(secp256k1CodeHash),
		Secp256k1CodeOutPoint:  secp256k1CodeOutPoint,
		Secp256k1DataOutPoint:  secp256k1DataOutPoint,
		Secp256k1GroupOutPoint: secp256k1GroupOutPoint,

		DaoOutPoint: daoOutPoint,
		DaoCodeHash: types.NewHexStr(daoCodeHash),
		DaoTypeHash: daoTypeHash,

		MultiSignSecpCellTypeHash:  multiSignTypeHash,
		MultiSignSecpGroupOutPoint: &multiSignSecpGroupOutPoint,
	}
	return sysCells, nil
}
