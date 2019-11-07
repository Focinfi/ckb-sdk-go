package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type WitnessArgs struct {
	WitnessesForInputLock  ByteFixVecOption
	WitnessesForInputType  ByteFixVecOption
	WitnessesForOutputType ByteFixVecOption
	serializer             *Table
}

func (wa WitnessArgs) Serialize() []byte {
	if wa.serializer != nil {
		return wa.serializer.Serialize()
	}
	return nil
}

func NewWitnessArgs(witness ckbtypes.Witness) (*WitnessArgs, error) {
	witnessesForInputLock, err := NewByteFixVecOption(witness.Lock)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidWitnessesForInputLock, nil)
	}
	witnessesForInputType, err := NewByteFixVecOption(witness.InputType)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidWitnessesForInputType, nil)
	}
	witnessesForOutputType, err := NewByteFixVecOption(witness.OutputType)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidWitnessesForOutputType, nil)
	}
	serializer := NewTable([]Serializer{
		witnessesForInputLock,
		witnessesForInputType,
		witnessesForOutputType,
	})
	return &WitnessArgs{
		WitnessesForInputLock:  *witnessesForInputLock,
		WitnessesForInputType:  *witnessesForInputType,
		WitnessesForOutputType: *witnessesForOutputType,
		serializer:             serializer,
	}, nil
}
