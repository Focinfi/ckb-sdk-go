package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type FixVec struct {
	ItemLen  Uint32
	ItemSize int // byte number
	data     []byte
}

func (fv FixVec) Serialize() []byte {
	return append(fv.Header(), fv.data...)
}

func (fv FixVec) Header() []byte {
	return fv.ItemLen.Serialize()
}

func NewFixVec(itemSize int, items []Serializer) (*FixVec, error) {
	data := make([]byte, 0)
	for _, item := range items {
		bs := item.Serialize()
		if len(bs) != itemSize {
			return nil, errtypes.WrapErr(errtypes.SerializationErrItemSizeNotFixed, nil)
		}
		data = append(data, bs...)
	}

	return &FixVec{
		ItemSize: itemSize,
		ItemLen:  Uint32(len(items)),
		data:     data,
	}, nil
}

type ByteFixVec struct {
	fixVec *FixVec
}

func (vec ByteFixVec) Serialize() []byte {
	return vec.fixVec.Serialize()
}

func NewByteFixVec(data []byte) *ByteFixVec {
	return &ByteFixVec{
		fixVec: &FixVec{
			ItemLen:  Uint32(len(data)),
			ItemSize: 1,
			data:     data,
		},
	}
}

func NewByteFixVecByHex(str string) (*ByteFixVec, error) {
	hexStr, err := types.ParseHexStr(str)
	if err != nil {
		return nil, err
	}
	return NewByteFixVec(hexStr.Bytes()), nil
}

type Byte32FixVec struct {
	serializer *FixVec
}

func NewByte32FixVec(items []Byte32) *ByteFixVec {
	data := make([]byte, 0, len(items)*32)
	for _, item := range items {
		data = append(data, item.Serialize()...)
	}
	return &ByteFixVec{
		fixVec: &FixVec{
			ItemLen:  Uint32(len(data)),
			ItemSize: 32,
			data:     data,
		},
	}
}

func NewByte32FixVecByHexes(hexes []string) (*ByteFixVec, error) {
	items := make([]Byte32, 0, len(hexes))
	for _, hex := range hexes {
		hexStr, err := types.ParseHexStr(hex)
		if err != nil {
			return nil, err
		}
		items = append(items, hexStr.Bytes())
	}
	return NewByte32FixVec(items), nil
}
