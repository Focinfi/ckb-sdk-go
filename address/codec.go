package address

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/bech32"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

// EncodeAddress encode the given prefix and payload data into the address with bech32
func EncodeAddress(prefix string, payload []byte) (string, error) {
	payload, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		return "", errtypes.WrapErr(errtypes.AddressErrConvertBitFail, err)
	}
	addr, err := bech32.Encode(prefix, payload)
	if err != nil {
		return "", errtypes.WrapErr(errtypes.CryptoErrBech32EncodeFail, err)
	}
	return addr, nil
}

// DecodeAddress decodes the given address into prefix and payload data
func DecodeAddress(address string) (prefix string, payload []byte, err error) {
	prefix, data, err := bech32.Decode(address)
	if err != nil {
		return "", nil, errtypes.WrapErr(errtypes.CryptoErrBech32DecodeFail, err)
	}
	if !addrtypes.IsAllowedPrefix(prefix) {
		return "", nil, errtypes.WrapErr(errtypes.AddressErrInvalidPrefix, nil)
	}
	payload, err = bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, errtypes.WrapErr(errtypes.AddressErrConvertBitFail, err)
	}
	return prefix, payload, nil
}
