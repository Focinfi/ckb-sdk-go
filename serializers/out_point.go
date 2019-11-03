package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type OutPoint struct {
	Index      OutPointIndex
	TxHash     OutPointTxHash
	serializer *Struct
}

func (outpoint OutPoint) Serialize() []byte {
	if outpoint.serializer != nil {
		return outpoint.serializer.Serialize()
	}
	return nil
}

func NewOutPoint(op ckbtypes.OutPoint) (*OutPoint, error) {
	index, err := NewUint32ByHex(op.Index)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidOutPointIndex, err)
	}
	txHash, err := NewByte32(op.TxHash)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidOutPointHash, err)
	}
	return &OutPoint{
		Index:  index,
		TxHash: txHash,
		serializer: &Struct{
			txHash,
			index,
		},
	}, nil
}
