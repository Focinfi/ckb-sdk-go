package address

import (
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

// Address represents a CKB address
// rfc-21: https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md
type Address struct {
	Mode    types.Mode
	KeyHash *KeyHash
	Prefix  string
}

// NewAddressFromPubKey creates and returns a new Address
func NewAddressFromPubKey(pubKey string, mode types.Mode) (*Address, error) {
	pk, err := NewPubKey(pubKey)
	if err != nil {
		return nil, err
	}
	return &Address{
		Mode:    mode,
		KeyHash: pk,
		Prefix:  addrtypes.PrefixForMode(mode),
	}, nil
}

// Generate generates address assuming default lock script is used
// payload = type(01) | code hash index(00) | pubkey Blake160
// see https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md for more info.
func (addr *Address) Generate() (string, error) {
	if addr.KeyHash == nil {
		return "", errtypes.WrapErr(errtypes.AddressErrEmptyPubKey, nil)
	}
	payload := append([]byte{
		byte(addrtypes.FormatTypeShort),
		byte(addrtypes.CodeHashIndex0)},
		addr.KeyHash.Blake160.Bytes()...)
	return EncodeAddress(addr.Prefix, payload)
}

// Generates short payload format address
// payload = type(01) | code hash index(01) | multisig
// see https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md for more info.
func (addr *Address) GenerateShortPayloadFormatAddress(hash160HexStr string) (string, error) {
	hexStr, err := types.ParseHexStr(hash160HexStr)
	if err != nil {
		return "", err
	}
	payload := append([]byte{
		byte(addrtypes.FormatTypeShort),
		byte(addrtypes.CodeHashIndex1)},
		hexStr.Bytes()...)
	return EncodeAddress(addr.Prefix, payload)
}

// GenerateFullPayloadAddress generates full payload format address
// payload = 0x02/0x04 | code_hash | args
// see https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md for more info.
func (addr *Address) GenerateFullPayloadAddress(formatType addrtypes.FormatType, codeHash, args string, mode types.Mode) (string, error) {
	if !addrtypes.IsFullPayloadFormatType(formatType) {
		return "", errtypes.WrapErr(errtypes.AddressErrFormatTypeWrong, nil)
	}
	prefix := addrtypes.PrefixForMode(mode)
	hexCodeHash, err := types.ParseHexStr(codeHash)
	if err != nil {
		return "", err
	}
	hexArgs, err := types.ParseHexStr(args)
	if err != nil {
		return "", err
	}

	payload := make([]byte, 0, 1+hexCodeHash.Len()+hexArgs.Len())
	payload = append(payload, byte(formatType))
	payload = append(payload, hexCodeHash.Bytes()...)
	payload = append(payload, hexArgs.Bytes()...)
	return EncodeAddress(prefix, payload)
}
