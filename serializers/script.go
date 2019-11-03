package serializers

import (
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

type Script struct {
	CodeHash   CodeHash
	HashType   HashType
	Args       ByteFixVec
	serializer *Table
}

func (script Script) Serialize() []byte {
	if script.serializer != nil {
		return script.serializer.Serialize()
	}
	return nil
}

func NewScript(script ckbtypes.Script) (*Script, error) {
	codeHash, err := NewByte32(script.CodeHash)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidCodeHash, err)
	}
	hashType, err := NewHashType(script.HashType)
	if err != nil {
		return nil, err
	}
	args, err := NewByteFixVecByHex(script.Args)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.SerializationErrInvalidArgs, err)
	}
	scr := Script{
		CodeHash: codeHash,
		HashType: *hashType,
		Args:     *args,
	}
	scr.serializer = NewTable([]Serializer{
		scr.CodeHash,
		scr.HashType,
		scr.Args,
	})
	return &scr, nil
}
