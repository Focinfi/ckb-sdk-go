package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type RawTransaction struct {
	Version    Version
	CellDeps   FixVec
	Header     ByteFixVec
	Inputs     FixVec
	Outputs    DynVec
	OutputData DynVec
	serializer *Table
}

func (t RawTransaction) Serialize() []byte {
	if t.serializer != nil {
		return t.serializer.Serialize()
	}
	return nil
}

func NewRawTransaction(transaction ckbtypes.Transaction) (*RawTransaction, error) {
	version, err := NewUint32ByHex(transaction.Version)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidVersion, err)
	}
	cellDeps, err := NewCellDepFixVec(transaction.CellDeps)
	if err != nil {
		return nil, err
	}
	headers, err := NewByte32FixVecByHexes(transaction.HeaderDeps)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidHeaderDep, err)
	}
	inputs, err := NewInputsFixVec(transaction.Inputs)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidHeaderDep, err)
	}
	output, err := NewOutputDynVec(transaction.Outputs)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidOutput, err)
	}
	outputData, err := NewHexDynVec(transaction.OutputsData)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidOutputData, err)
	}
	serializer := NewTable([]Serializer{
		version, cellDeps, headers, inputs, output, outputData,
	})
	return &RawTransaction{
		Version:    version,
		CellDeps:   *cellDeps,
		Header:     *headers,
		Inputs:     *inputs,
		Outputs:    *output,
		OutputData: *outputData,
		serializer: serializer,
	}, nil
}
