package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Output struct {
	Capacity   Capacity
	LockScript Script
	TypeScript *Script
	serializer *Table
}

func (output *Output) InitSerializer() {
	fields := []Serializer{
		output.Capacity,
		output.LockScript,
	}
	if output.TypeScript != nil {
		fields = append(fields, output.TypeScript)
	} else {
		fields = append(fields, Empty{})
	}
	output.serializer = NewTable(fields)
}

func (output Output) Serialize() []byte {
	if output.serializer != nil {
		return output.serializer.Serialize()
	}
	return nil
}

func NewOutput(output ckbtypes.Output) (*Output, error) {
	capacity, err := NewUint64ByHex(output.Capacity)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrOutputCapacity, err)
	}
	lockScript, err := NewScript(output.Lock)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidLockScript, err)
	}
	var typeScript *Script
	if output.Type != nil {
		typeScript, err = NewScript(output.Lock)
		if err != nil {
			return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidTypeScript, err)
		}
	}
	op := &Output{
		Capacity:   capacity,
		LockScript: *lockScript,
		TypeScript: typeScript,
	}
	op.InitSerializer()
	return op, nil
}

func NewOutputsDynVec(outputs []ckbtypes.Output) (*DynVec, error) {
	items := make([]Serializer, 0, len(outputs))
	for _, output := range outputs {
		op, err := NewOutput(output)
		if err != nil {
			return nil, err
		}
		items = append(items, op)
	}
	return NewDynVec(items), nil
}
