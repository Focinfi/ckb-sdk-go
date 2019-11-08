package utils

import (
	"errors"

	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"
	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/serializers"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

const emptiedWitnessLock = "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

func SignTransaction(key key.Key, transaction ckbtypes.Transaction) (*ckbtypes.Transaction, error) {
	if len(transaction.Witnesses) == 0 {
		return nil, errtypes.WrapErr(errtypes.GenTransErrWitnessNotEnough, nil)
	}
	firstWitness, ok := transaction.Witnesses[0].(ckbtypes.Witness)
	if !ok {
		return nil, errtypes.WrapErr(errtypes.GenTransErrFirstWitnessTypeWrong, errors.New("first witness must be a ckbtypes.Witness"))
	}
	txHash, err := RawTransactionHash(transaction)
	if err != nil {
		return nil, err
	}
	emptiedWitness := firstWitness.Clone()
	emptiedWitness.Lock = emptiedWitnessLock
	emptiedWitnessSerializer, err := serializers.NewWitnessArgs(*emptiedWitness)
	if err != nil {
		return nil, err
	}
	emptiedWitnessData := emptiedWitnessSerializer.Serialize()
	emptiedWitnessDataSize := len(emptiedWitnessData)

	data := append(txHash.Bytes(), types.HexUint64(emptiedWitnessDataSize).LittleEndianBytes(8)...)
	data = append(data, emptiedWitnessData...)
	for i := 1; i < len(transaction.Witnesses); i++ {
		switch w := transaction.Witnesses[i].(type) {
		case ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(w)
			if err != nil {
				return nil, err
			}
			data = append(data, slr.Serialize()...)
		case *ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(*w)
			if err != nil {
				return nil, err
			}
			data = append(data, slr.Serialize()...)
		case string:
			hexStr, err := types.ParseHexStr(w)
			if err != nil {
				return nil, err
			}
			data = append(data, hexStr.Bytes()...)
		default:
			return nil, errtypes.WrapErr(errtypes.GenTransErrHexWitnessTypeWrong, nil)
		}
	}
	message, err := blake2b.Digest(data)
	if err != nil {
		return nil, err
	}
	sign, err := key.SignRecoverableFor32Bytes(message)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.GenTransErrSignFail, err)
	}
	emptiedWitness.Lock = sign.Hex()

	// renew the serializer
	emptiedWitnessSerializer, err = serializers.NewWitnessArgs(*emptiedWitness)
	if err != nil {
		return nil, err
	}
	emptiedWitnessHex := types.NewHexStr(emptiedWitnessSerializer.Serialize()).Hex()

	witness := []interface{}{emptiedWitnessHex}
	if len(transaction.Witnesses) > 1 {
		witness = append(witness, transaction.Witnesses...)
	}
	return &ckbtypes.Transaction{
		Hash:        txHash.Hex(),
		Version:     transaction.Version,
		CellDeps:    transaction.CellDeps,
		HeaderDeps:  transaction.HeaderDeps,
		Inputs:      transaction.Inputs,
		Outputs:     transaction.Outputs,
		OutputsData: transaction.OutputsData,
		Witnesses:   witness,
	}, nil
}
