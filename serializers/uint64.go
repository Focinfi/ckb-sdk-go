package serializers

import "github.com/Focinfi/ckb-sdk-go/types"

type Uint64 uint64

// Serialize serializes in little endian
func (h Uint64) Serialize() []byte {
	return types.HexUint64(uint64(h)).BigEndianBytes(8)
}

func NewUint64ByHex(hexStr string) (Uint64, error) {
	n, err := types.ParseHexUint64(hexStr)
	if err != nil {
		return 0, err
	}
	return Uint64(n.Uint64()), nil
}
