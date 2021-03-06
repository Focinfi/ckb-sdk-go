package key

import (
	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const privKeyCharCount = 66

// Key represents a asymmetric encryption pair of private key and public key
type Key struct {
	privKeyHex *types.HexStr
	pubKeyHex  *types.HexStr
	privKey    *btcec.PrivateKey
	pubKey     *btcec.PublicKey
	Address    *address.Address
}

// NewFromPrivKeyHex creates and returns a new Key depends on the given private 0x prefixed hex string and mode using secp256k1
func NewFromPrivKeyHex(hexStr string, mode types.Mode) (*Key, error) {
	if len(hexStr) != privKeyCharCount {
		return nil, errtypes.WrapErr(errtypes.KeyErrPrivateKeySizeWrong, nil)
	}
	privKeyHex, err := types.ParseHexStr(hexStr)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, err)
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyHex.Bytes())
	pubKeyHex := types.NewHexStr(pubKey.SerializeCompressed())
	addr, err := address.NewAddressFromPubKey(pubKeyHex.Hex(), mode)
	if err != nil {
		return nil, err
	}
	return &Key{
		privKeyHex: privKeyHex,
		pubKeyHex:  pubKeyHex,
		privKey:    privKey,
		pubKey:     pubKey,
		Address:    addr,
	}, nil
}

// Sign signs data of decoding from the given hexStr using the private key
func (key *Key) Sign(hexStr string) (*types.HexStr, error) {
	dataHex, err := types.ParseHexStr(hexStr)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, err)
	}
	sign, err := key.privKey.Sign(dataHex.Bytes())
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.CryptoErrSignDataFail, err)
	}

	return types.NewHexStr(sign.Serialize()), nil
}

// SignRecoverableFor32BytesHex signs the data of decoding from the given hex32BytesStr using the private key in recoverable mode
func (key *Key) SignRecoverableFor32BytesHex(hex32BytesStr string) (*types.HexStr, error) {
	dataHex, err := types.ParseHexStr(hex32BytesStr)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, err)
	}
	return key.SignRecoverableFor32Bytes(dataHex.Bytes())
}

func (key *Key) SignRecoverableFor32Bytes(data []byte) (*types.HexStr, error) {
	if len(data) != 32 {
		return nil, errtypes.WrapErr(errtypes.CryptoErrDataByteCountNot32, nil)
	}
	sign, err := secp256k1.Sign(data, key.privKeyHex.Bytes())
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.CryptoErrSignRecoverableFail, err)
	}
	return types.NewHexStr(sign), nil
}
