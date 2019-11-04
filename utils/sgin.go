package utils

import (
	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

func SignTransaction(key key.Key, transaction *ckbtypes.Transaction) error {
	if transaction == nil {
		return nil
	}
	txHash, err := TransactionHash(*transaction)
	if err != nil {
		return err
	}
	if len(transaction.Witnesses) < len(transaction.Inputs) {
		return errtypes.WrapErr(errtypes.GenTransErrWitnessNumLessThanInputs, nil)
	}
	signedWitnesses := make([]string, 0, len(transaction.Witnesses))
	for _, w := range transaction.Witnesses {
		old, err := types.ParseHexStr(w)
		if err != nil {
			return nil
		}

		data := old.Append(*txHash)
		message, err := blake2b.Digest(data.Bytes())
		if err != nil {
			return errtypes.WrapErr(errtypes.GenTransErrSignFail, err)
		}
		sign, err := key.SignRecoverableFor32Bytes(message)
		if err != nil {
			return errtypes.WrapErr(errtypes.GenTransErrSignFail, err)
		}
		witness := sign.Append(*old)
		signedWitnesses = append(signedWitnesses, witness.Hex())
	}
	transaction.Hash = txHash.Hex()
	transaction.Witnesses = signedWitnesses
	return nil
}
