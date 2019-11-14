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
		return nil, errtypes.WrapErr(errtypes.GenTransErrFirstWitnessTypeWrong, errors.New("first witnesses must be a ckbtypes.Witness"))
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

	toSignData := append(txHash.Bytes(), types.HexUint64(emptiedWitnessDataSize).LittleEndianBytes(8)...)
	toSignData = append(toSignData, emptiedWitnessData...)
	for i := 1; i < len(transaction.Witnesses); i++ {
		var witnessData []byte
		switch w := transaction.Witnesses[i].(type) {
		case ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(w)
			if err != nil {
				return nil, err
			}
			witnessData = slr.Serialize()
		case *ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(*w)
			if err != nil {
				return nil, err
			}
			witnessData = slr.Serialize()
		case string:
			hexStr, err := types.ParseHexStr(w)
			if err != nil {
				return nil, err
			}

			witnessData = hexStr.Bytes()
		default:
			return nil, errtypes.WrapErr(errtypes.GenTransErrHexWitnessTypeWrong, nil)
		}
		witnessDataSize := types.HexUint64(len(witnessData))
		toSignData = append(toSignData, witnessDataSize.LittleEndianBytes(8)...)
		toSignData = append(toSignData, witnessData...)
	}
	message, err := blake2b.Digest(toSignData)
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

	witnesses := []interface{}{emptiedWitnessHex}
	if len(transaction.Witnesses) > 1 {
		witnesses = append(witnesses, transaction.Witnesses[1:]...)
	}
	return &ckbtypes.Transaction{
		Hash:        txHash.Hex(),
		Version:     transaction.Version,
		CellDeps:    transaction.CellDeps,
		HeaderDeps:  transaction.HeaderDeps,
		Inputs:      transaction.Inputs,
		Outputs:     transaction.Outputs,
		OutputsData: transaction.OutputsData,
		Witnesses:   witnesses,
	}, nil
}

func MultiSignTransaction(privKeys []string, transaction ckbtypes.Transaction, configSerializationHex *types.HexStr, mode types.Mode) (*ckbtypes.Transaction, error) {
	if len(transaction.Witnesses) == 0 {
		return nil, errtypes.WrapErr(errtypes.GenTransErrWitnessNotEnough, nil)
	}
	firstWitness, ok := transaction.Witnesses[0].(ckbtypes.Witness)
	if !ok {
		return nil, errtypes.WrapErr(errtypes.GenTransErrFirstWitnessTypeWrong, errors.New("first witnesses must be a ckbtypes.Witness"))
	}
	txHash, err := RawTransactionHash(transaction)
	if err != nil {
		return nil, err
	}

	emptiedWitness := firstWitness.Clone()
	emptySignature := make([]byte, 65*len(privKeys))
	emptiedWitness.Lock = configSerializationHex.AppendBytes(emptySignature).Hex()
	emptiedWitnessDataSlr, err := serializers.NewWitnessArgs(*emptiedWitness)
	if err != nil {
		return nil, err
	}
	emptiedWitnessData := emptiedWitnessDataSlr.Serialize()
	emptiedWitnessDataSize := types.HexUint64(len(emptiedWitnessData))

	toSignData := append(txHash.Bytes(), emptiedWitnessDataSize.LittleEndianBytes(8)...)
	toSignData = append(toSignData, emptiedWitnessData...)

	for i := 1; i < len(transaction.Witnesses); i++ {
		var witnessData []byte
		switch w := transaction.Witnesses[i].(type) {
		case ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(w)
			if err != nil {
				return nil, err
			}
			witnessData = slr.Serialize()
		case *ckbtypes.Witness:
			slr, err := serializers.NewWitnessArgs(*w)
			if err != nil {
				return nil, err
			}
			witnessData = slr.Serialize()
		case string:
			hexStr, err := types.ParseHexStr(w)
			if err != nil {
				return nil, err
			}
			witnessData = hexStr.Bytes()
		}

		witnessDataSize := types.HexUint64(len(witnessData))
		toSignData = append(toSignData, witnessDataSize.LittleEndianBytes(8)...)
		toSignData = append(toSignData, witnessData...)
	}
	message, err := blake2b.Digest(toSignData)
	if err != nil {
		return nil, err
	}
	signatures := make([]byte, 0, len(privKeys)*65)
	for _, privKey := range privKeys {
		k, err := key.NewFromPrivKeyHex(privKey, mode)
		if err != nil {
			return nil, err
		}
		h, err := k.SignRecoverableFor32Bytes(message)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, h.Bytes()...)
	}
	lockHex := configSerializationHex.AppendBytes(signatures)
	emptiedWitness.Lock = lockHex.Hex()
	emptiedWitnessSlr, err := serializers.NewWitnessArgs(*emptiedWitness)
	if err != nil {
		return nil, err
	}
	witnesses := []interface{}{types.NewHexStr(emptiedWitnessSlr.Serialize()).Hex()}
	if len(transaction.Witnesses) > 1 {
		witnesses = append(witnesses, transaction.Witnesses[1:]...)
	}

	return &ckbtypes.Transaction{
		Hash:        txHash.Hex(),
		Version:     transaction.Version,
		CellDeps:    transaction.CellDeps,
		HeaderDeps:  transaction.HeaderDeps,
		Inputs:      transaction.Inputs,
		Outputs:     transaction.Outputs,
		OutputsData: transaction.OutputsData,
		Witnesses:   witnesses,
	}, nil
}
