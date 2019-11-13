package blake2b

import (
	"github.com/dchest/blake2b"
)

const personal = "ckb-default-hash"

// Digest digests the data with blake2b
func Digest(data []byte) ([]byte, error) {
	config := &blake2b.Config{Size: 32, Person: []byte(personal)}
	h, err := blake2b.New(config)
	if err != nil {
		return nil, err
	}
	h.Write(data)
	return h.Sum(nil), nil
}
