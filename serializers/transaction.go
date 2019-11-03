package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Transaction struct {
	RawTransaction RawTransaction
	Witness        ByteDynVec
	serializer     *Table
}

func NewTransaction(transaction ckbtypes.Transaction) (*Transaction, error) {
	rawTrans, err := NewRawTransaction(transaction)
	if err != nil {
		return nil, err
	}
	witness, err := NewByteDynVecByHexes(transaction.Witnesses)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidWitness, err)
	}
	seriealizer := NewTable([]Serializer{rawTrans, witness})
	return &Transaction{
		RawTransaction: *rawTrans,
		Witness:        *witness,
		serializer:     seriealizer,
	}, nil
}
