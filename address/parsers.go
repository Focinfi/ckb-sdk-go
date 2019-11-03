package address

import (
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func Parse(address string, mode types.Mode) ([]string, error) {
	prefix, payload, err := decodeAddress(address)
	if err != nil {
		return nil, err
	}
	if prefix != addrtypes.PrefixForMode(mode) {
		return nil, errtypes.WrapErr(errtypes.AddressErrInvalidPrefix, nil)
	}
	if len(payload) <= 1 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}

	formatType := addrtypes.FormatType(payload[0])
	switch formatType {
	case addrtypes.FormatTypeShortLock:
		return parseShortPayloadAddress(formatType, payload[1:])
	case addrtypes.FormatTypeCode, addrtypes.FormatTypeData:
		return parseFullPayloadAddress(formatType, payload[1:])
	}

	return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
}

// ParseShortPayloadAddress
// address = ckt/ckb | 0x01 | code_hash_index | single_arg
// return = [hex(code_hash_index), hex(single_arg)]
func ParseShortPayloadAddress(address string, mode types.Mode) ([]string, error) {
	prefix, payload, err := decodeAddress(address)
	if err != nil {
		return nil, err
	}
	if prefix != addrtypes.PrefixForMode(mode) {
		return nil, errtypes.WrapErr(errtypes.AddressErrInvalidPrefix, nil)
	}
	if len(payload) <= 1 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}
	formatType := addrtypes.FormatType(payload[0])
	return parseShortPayloadAddress(formatType, payload[1:])
}

func parseShortPayloadAddress(formatType addrtypes.FormatType, payload []byte) ([]string, error) {
	if formatType != addrtypes.FormatTypeShortLock {
		return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
	}
	hashInfo, err := addrtypes.NewHashInfo(payload)
	if err != nil {
		return nil, err
	}
	return []string{formatType.Hex(), hashInfo.HashType.Hex(), hashInfo.Data.Hex()}, nil
}

// ParseFullPayloadAddress
// payload = ckt/ckb | 0x02/0x04(1bit) | code_hash(32bit) | args
// return = ["0x2"/"0x4", hex(code_hash), hex(args)]
func ParseFullPayloadAddress(address string, mode types.Mode) ([]string, error) {
	prefix, payload, err := decodeAddress(address)
	if err != nil {
		return nil, err
	}
	if prefix != addrtypes.PrefixForMode(mode) {
		return nil, errtypes.WrapErr(errtypes.AddressErrInvalidPrefix, nil)
	}
	if len(payload) <= 1 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}
	formatType := addrtypes.FormatType(payload[0])
	return parseFullPayloadAddress(formatType, payload[1:])
}

func parseFullPayloadAddress(formatType addrtypes.FormatType, payload []byte) ([]string, error) {
	if addrtypes.IsFullPayloadFormatType(formatType) {
		return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
	}
	if len(payload) <= 32 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}
	return []string{
		formatType.Hex(),
		types.NewHexStr(payload[:32]).Hex(),
		types.NewHexStr(payload[32:]).Hex(),
	}, nil
}
