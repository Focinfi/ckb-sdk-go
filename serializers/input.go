package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Input struct {
	PreviousOutput OutPoint
	Since          Since
	serializer     *Struct
}

func (input Input) Serialize() []byte {
	if input.serializer != nil {
		return input.serializer.Serialize()
	}
	return nil
}

func NewInput(input ckbtypes.Input) (*Input, error) {
	prvOutput, err := NewOutPoint(input.PreviousOutput)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidInput, err)
	}
	since, err := NewUint64ByHex(input.Since)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidSince, err)
	}
	return &Input{
		PreviousOutput: *prvOutput,
		Since:          since,
		serializer: &Struct{
			*prvOutput,
			since,
		},
	}, nil
}

func NewInputsFixVec(inputs []ckbtypes.Input) (*FixVec, error) {
	items := make([]Serializer, 0, len(inputs))
	for _, in := range inputs {
		input, err := NewInput(in)
		if err != nil {
			return nil, err
		}
		items = append(items, input)
	}
	vec, err := NewFixVec(InputCapacity, items)
	if err != nil {
		return nil, err
	}
	return vec, nil
}
