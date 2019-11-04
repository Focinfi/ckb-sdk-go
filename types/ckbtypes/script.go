package ckbtypes

import "github.com/Focinfi/ckb-sdk-go/types"

type Script struct {
	Args     string   `json:"args"`
	CodeHash string   `json:"code_hash"`
	HashType HashType `json:"hash_type"`
}

func (script Script) ByteSize() (uint64, error) {
	size := uint64(1) // HashType(1 byte)
	if len(script.CodeHash) > 2 {
		hexStr, err := types.ParseHexStr(script.CodeHash)
		if err != nil {
			return 0, nil
		}
		size += uint64(len(hexStr.Bytes()))
	}
	if len(script.Args) > 2 {
		hexStr, err := types.ParseHexStr(script.Args)
		if err != nil {
			return 0, nil
		}
		size += uint64(len(hexStr.Bytes()))
	}
	return size, nil
}
