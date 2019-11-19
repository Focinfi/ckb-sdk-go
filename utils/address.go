package utils

import (
	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func LockScriptFormAddress(addr string, mode types.Mode, sysCells SysCells) (*ckbtypes.Script, error) {
	config, err := address.Parse(addr, mode)
	if err != nil {
		return nil, err
	}
	var codeHash *types.HexStr
	switch config.FormatType {
	case addrtypes.FormatTypeShort:
		switch config.CodeHashIndex {
		case addrtypes.CodeHashIndex0:
			codeHash = sysCells.Secp256k1TypeHash
		case addrtypes.CodeHashIndex1:
			codeHash = sysCells.MultiSignSecpCellTypeHash
		}
	case addrtypes.FormatTypeFullType, addrtypes.FormatTypeFullData:
		codeHash = config.CodeHash
	}

	if codeHash == nil {
		return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
	}

	return &ckbtypes.Script{
		Args:     config.Args.Hex(),
		CodeHash: codeHash.Hex(),
		HashType: ckbtypes.HashTypeType,
	}, nil
}
