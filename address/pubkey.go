package address

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/blake160"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

// KeyHash contains the origin public key in hex 0x prefixed string and its blake160-hash hex string
type KeyHash struct {
	PubKey   string
	Blake160 *types.HexStr
}

// NewPubKey creates and returns a new KeyHash.
// Assumes that the given pubKey is the 0x prefixed hex number string
func NewPubKey(pubKey string) (*KeyHash, error) {
	blake160hash, err := blake160.Blake160(pubKey)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.CryptoErrBlake160Fail, err)
	}
	return &KeyHash{
		PubKey:   pubKey,
		Blake160: types.NewHexStr(blake160hash),
	}, nil
}
