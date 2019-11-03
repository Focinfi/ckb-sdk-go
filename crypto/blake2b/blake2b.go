package blake2b

import (
	"encoding/hex"

	"github.com/dchest/blake2b"
)

const personal = "ckb-default-hash"

func Digest(data []byte) ([]byte, error) {
	config := &blake2b.Config{Size: 32, Person: []byte(personal)}
	h, err := blake2b.New(config)
	if err != nil {
		return nil, err
	}
	h.Write(data)
	return h.Sum(nil), nil
}

func HexDigest(data []byte) (string, error) {
	b, err := Digest(data)
	if err != nil {
		return "", err
	}
	return "0x" + hex.EncodeToString(b), nil
}
