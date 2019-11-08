package utils

import (
	"github.com/Focinfi/ckb-sdk-go/types"
)

type Epoch struct {
	Length uint64
	Number uint64
	Index  uint64
}

func (e Epoch) Since() uint64 {
	return (0x20 << 56) + (e.Length << 40) + (e.Index << 24) + e.Number
}

func ParseEpochByHexStr(hexStr string) (*Epoch, error) {
	hexNum, err := types.ParseHexUint64(hexStr)
	if err != nil {
		return nil, err
	}
	num := hexNum.Uint64()
	return &Epoch{
		Length: (num >> 40) & 0xFFFF,
		Index:  (num >> 24) & 0xFFFF,
		Number: num & 0xFFFFFF,
	}, nil
}
