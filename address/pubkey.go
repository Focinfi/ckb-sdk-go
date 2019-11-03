package address

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/blake160"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type PubKey struct {
	PubKey   string
	Blake160 *types.HexStr
}

func NewPubKey(pubKey string) (*PubKey, error) {
	blake160hash, err := blake160.Blake160(pubKey)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.CryptoErrBlake160Fail, err)
	}
	return &PubKey{
		PubKey:   pubKey,
		Blake160: types.NewHexStr(blake160hash),
	}, nil
}
