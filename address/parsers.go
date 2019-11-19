package address

import (
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

// AddrConfig contains the parts a address
type AddrConfig struct {
	FormatType    addrtypes.FormatType
	CodeHashIndex addrtypes.CodeHashIndex
	CodeHash      *types.HexStr
	Args          *types.HexStr
}

// Parse parses the given address in the given mode, returns a AddressConfig
func Parse(address string, mode types.Mode) (*AddrConfig, error) {
	prefix, payload, err := DecodeAddress(address)
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
	case addrtypes.FormatTypeShort:
		return parseShortPayloadAddress(formatType, payload[1:])
	case addrtypes.FormatTypeFullType, addrtypes.FormatTypeFullData:
		return parseFullPayloadAddress(formatType, payload[1:])
	}

	return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, fmt.Errorf("unsupported format type: %d", formatType))
}

// ParseShortPayloadAddress
// address = ckt/ckb | 0x01 | code_hash_index | single_arg
// return = [0x01, hex(code_hash_index), hex(single_arg)]
func ParseShortPayloadAddress(address string, mode types.Mode) (*AddrConfig, error) {
	prefix, payload, err := DecodeAddress(address)
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

func ParseShortPayloadAddressArg(address string, mode types.Mode) (*types.HexStr, error) {
	config, err := ParseShortPayloadAddress(address, mode)
	if err != nil {
		return nil, err
	}
	return config.Args, nil
}

func parseShortPayloadAddress(formatType addrtypes.FormatType, payload []byte) (*AddrConfig, error) {
	if formatType != addrtypes.FormatTypeShort {
		return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
	}
	hashInfo, err := addrtypes.NewHashInfo(payload)
	if err != nil {
		return nil, err
	}
	return &AddrConfig{
		FormatType:    formatType,
		CodeHashIndex: hashInfo.HashType,
		Args:          hashInfo.Data,
	}, nil
}

// ParseFullPayloadAddress
// payload = ckt/ckb | 0x02/0x04(1bit) | code_hash(32bit) | args
// return = ["0x2"/"0x4", hex(code_hash), hex(args)]
func ParseFullPayloadAddress(address string, mode types.Mode) (*AddrConfig, error) {
	prefix, payload, err := DecodeAddress(address)
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

func parseFullPayloadAddress(formatType addrtypes.FormatType, payload []byte) (*AddrConfig, error) {
	if !addrtypes.IsFullPayloadFormatType(formatType) {
		return nil, errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, fmt.Errorf("format type (0x%x) is not 0x2 or 0x4", formatType))
	}
	if len(payload) <= 32 {
		return nil, errtypes.WrapErr(errtypes.AddressErrTooShort, nil)
	}
	return &AddrConfig{
		FormatType: formatType,
		CodeHash:   types.NewHexStr(payload[:32]),
		Args:       types.NewHexStr(payload[32:]),
	}, nil
}
