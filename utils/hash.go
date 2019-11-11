package utils

import (
	"fmt"

	"github.com/Focinfi/ckb-sdk-go/address"
	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/serializers"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func ScriptHash(script ckbtypes.Script) (*types.HexStr, error) {
	slr, err := serializers.NewScript(script)
	if err != nil {
		return nil, err
	}
	hash, err := blake2b.Digest(slr.Serialize())
	if err != nil {
		return nil, err
	}
	return types.NewHexStr(hash), nil
}

func NewLockScript(pubKey string) (*ckbtypes.Script, error) {
	k, err := address.NewPubKey(pubKey)
	if err != nil {
		return nil, err
	}
	fmt.Println("lock args:", k.Blake160.Hex())
	return &ckbtypes.Script{
		Args:     k.Blake160.Hex(),
		CodeHash: types.BlockAssemblerCodeHash,
		HashType: ckbtypes.HashTypeType,
	}, nil
}

func LockScriptHash(pubKey string) (*types.HexStr, error) {
	script, err := NewLockScript(pubKey)
	if err != nil {
		return nil, err
	}
	return ScriptHash(*script)
}

func RawTransactionHash(transaction ckbtypes.Transaction) (*types.HexStr, error) {
	slr, err := serializers.NewRawTransaction(transaction)
	if err != nil {
		return nil, err
	}
	hash, err := blake2b.Digest(slr.Serialize())
	if err != nil {
		return nil, err
	}
	return types.NewHexStr(hash), nil
}
