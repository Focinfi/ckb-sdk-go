package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Transaction struct {
	RawTransaction RawTransaction
	Witnesses      Witnesses
	serializer     *Table
}

func (transaction *Transaction) Serialize() []byte {
	if transaction.serializer != nil {
		return transaction.serializer.Serialize()
	}
	return nil
}

func NewTransaction(transaction ckbtypes.Transaction) (*Transaction, error) {
	rawTrans, err := NewRawTransaction(transaction)
	if err != nil {
		return nil, err
	}
	witnesses, err := NewWitnessesByHexes(transaction.Witnesses)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidWitness, err)
	}
	serializer := NewTable([]Serializer{rawTrans, witnesses})
	return &Transaction{
		RawTransaction: *rawTrans,
		Witnesses:      *witnesses,
		serializer:     serializer,
	}, nil
}
