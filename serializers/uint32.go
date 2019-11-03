package serializers

import "github.com/Focinfi/ckb-sdk-go/types"

type Uint32 uint32

// Serialize serializes in little endian
func (h Uint32) Serialize() []byte {
	return types.HexUint64(uint64(h)).LittleEndianBytes(4)
}

func NewUint32ByHex(hexStr string) (Uint32, error) {
	n, err := types.ParseHexUint64(hexStr)
	if err != nil {
		return 0, err
	}
	return Uint32(n.Uint64()), nil
}
