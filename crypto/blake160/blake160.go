package blake160

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/types"
)

func Blake160(pubKey string) ([]byte, error) {
	hexStr, err := types.ParseHexStr(pubKey)
	if err != nil {
		return nil, err
	}

	d, err := blake2b.Digest(hexStr.Bytes())
	if err != nil {
		return nil, err
	}
	return d[:20], nil
}
