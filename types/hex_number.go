package types

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type HexUint64 uint64

func (hexNumber HexUint64) Uint64() uint64 {
	return uint64(hexNumber)
}

func (hexNumber *HexUint64) UnmarshalJSON(data []byte) error {
	var ns string
	if err := json.Unmarshal(data, &ns); err != nil {
		return err
	}
	hexStr, err := ParseHexStr(ns)
	if err != nil {
		return err
	}
	n, err := hexStr.ToHexUint64()
	if err != nil {
		return err
	}
	*hexNumber = HexUint64(n)
	return nil
}

func (hexNumber HexUint64) Hex() string {
	return fmt.Sprintf("0x%x", hexNumber)
}

func (hexNumber HexUint64) LittleEndianString(paddingBytes int) string {
	bs := hexNumber.LittleEndianBytes(paddingBytes)
	return hex.EncodeToString(bs)
}

func (hexNumber HexUint64) LittleEndianBytes(paddingBytes int) []byte {
	switch paddingBytes {
	case 2:
		bs := [2]byte{}
		binary.LittleEndian.PutUint16(bs[:], uint16(hexNumber))
		return bs[:]
	case 4:
		bs := [4]byte{}
		binary.LittleEndian.PutUint32(bs[:], uint32(hexNumber))
		return bs[:]
	case 8:
		bs := [8]byte{}
		binary.LittleEndian.PutUint64(bs[:], uint64(hexNumber))
		return bs[:]
	}
	return nil
}

func (hexNumber HexUint64) BigEndianBytes(paddingBytes int) []byte {
	switch paddingBytes {
	case 2:
		bs := [2]byte{}
		binary.BigEndian.PutUint16(bs[:], uint16(hexNumber))
		return bs[:]
	case 4:
		bs := [4]byte{}
		binary.BigEndian.PutUint32(bs[:], uint32(hexNumber))
		return bs[:]
	case 8:
		bs := [8]byte{}
		binary.BigEndian.PutUint64(bs[:], uint64(hexNumber))
		return bs[:]
	}
	return nil
}

func ParseHexUint64(str string) (*HexUint64, error) {
	hexStr, err := ParseHexStr(str)
	if err != nil {
		return nil, err
	}
	n, err := strconv.ParseUint(hexStr.Hex()[2:], 16, 64)
	if err != nil {
		return nil, errors.New("parse hex number err: " + err.Error())
	}
	hexNumber := HexUint64(n)
	return &hexNumber, nil
}
