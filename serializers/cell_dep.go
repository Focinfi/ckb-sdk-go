package serializers

import (
	"errors"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type CellDep struct {
	OutPoint   OutPoint
	DepType    DepType
	serializer *Struct
}

func (cd CellDep) Serialize() []byte {
	if cd.serializer != nil {
		return cd.serializer.Serialize()
	}
	return nil
}

func NewCellDepFixVec(cds []ckbtypes.CellDep) (*FixVec, error) {
	cellDeps := make([]Serializer, 0, len(cds))
	for _, cd := range cds {
		outpoint, err := NewOutPoint(cd.OutPoint)
		if err != nil {
			return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidCellDep, err)
		}
		depType, err := NewDepType(cd.DepType)
		if err != nil {
			return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidCellDep, err)
		}
		cellDeps = append(cellDeps, CellDep{
			OutPoint:   *outpoint,
			DepType:    *depType,
			serializer: &Struct{*outpoint, *depType},
		})
	}
	vec, err := NewFixVec(OutPointCapacity, cellDeps)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidCellDep, err)
	}
	return vec, nil
}

func NewDepType(dt ckbtypes.DepType) (*DepType, error) {
	switch dt {
	case ckbtypes.DepTypeCode:
		return &DepTypeCode, nil
	case ckbtypes.DepTypeDepGroup:
		return &DepTypeDepGroup, nil
	}
	return nil, errtypes.WrapErr(errtypes.SerializationErrUnknownDepType, errors.New(string(dt)))
}
