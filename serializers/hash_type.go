package serializers

import (
	"errors"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func NewHashType(hashType ckbtypes.HashType) (*HashType, error) {
	switch hashType {
	case ckbtypes.HashTypeType:
		return &HashTypeType, nil
	case ckbtypes.HashTypeData:
		return &HashTypeData, nil
	}
	return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidHashType, errors.New(string(hashType)))
}
