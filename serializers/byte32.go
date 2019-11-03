package serializers

import (
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Byte32 []byte

func (b Byte32) Serialize() []byte {
	return b[:32]
}

func NewByte32(hex string) (Byte32, error) {
	hexStr, err := types.ParseHexStr(hex)
	if err != nil {
		return nil, err
	}
	if hexStr.Len() != 32 {
		return nil, errtypes.WrapErr(errtypes.SerializationErrByte32WrongLen, fmt.Errorf("len: %d", hexStr.Len()))
	}
	return Byte32(hexStr.Bytes()[:32]), nil
}
