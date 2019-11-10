package addrtypes

import (
	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type CodeHashIndex = uint8

// Code hash index
// doc: https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md#short-payload-format
const (
	CodeHashIndex0 CodeHashIndex = 0
	CodeHashIndex1 CodeHashIndex = 1
)

var codeHashIndexList = []CodeHashIndex{CodeHashIndex0, CodeHashIndex1}

func IsAllowedHashType(hashType CodeHashIndex) bool {
	for _, ht := range codeHashIndexList {
		if hashType == ht {
			return true
		}
	}
	return false
}

type HashInfo struct {
	HashType CodeHashIndex
	Data     *types.HexStr
}

func NewHashInfo(payload []byte) (*HashInfo, error) {
	if len(payload) == 0 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}
	hashType := CodeHashIndex(payload[0])
	if !IsAllowedHashType(hashType) {
		return nil, errtypes.WrapErr(errtypes.AddressErrInvalidHashTypeIndex, nil)
	}
	return &HashInfo{
		HashType: hashType,
		Data:     types.NewHexStr(payload[1:]),
	}, nil
}
