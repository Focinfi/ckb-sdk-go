package blake160

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/types"
)

// Blake160 decodes the pubKey as hex 0x prefixed string, returns its digest
func Blake160(pubKey string) ([]byte, error) {
	hexStr, err := types.ParseHexStr(pubKey)
	if err != nil {
		return nil, err
	}

	return Blake160Binary(hexStr.Bytes())
}

// Blake160Binary digests the given bin data in blake2b, returns first 20 bytes
func Blake160Binary(bin []byte) ([]byte, error) {
	d, err := blake2b.Digest(bin)
	if err != nil {
		return nil, err
	}
	return d[:20], nil
}
