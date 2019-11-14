package ckbtypes

import "github.com/Focinfi/ckb-sdk-go/types/errtypes"

type Output struct {
	Capacity string  `json:"capacity"`
	Lock     Script  `json:"lock"`
	Type     *Script `json:"type,omitempty"`
}

func (output Output) Clone() *Output {
	c := &Output{
		Capacity: output.Capacity,
		Lock: Script{
			Args:     output.Lock.Args,
			CodeHash: output.Lock.CodeHash,
			HashType: output.Lock.HashType,
		},
	}
	if output.Type != nil {
		c.Type = &Script{
			Args:     output.Type.Args,
			CodeHash: output.Type.CodeHash,
			HashType: output.Type.HashType,
		}
	}
	return c
}

func (output Output) ByteSize() (uint64, error) {
	size := uint64(8) // capacity(32 bit)
	if output.Type != nil {
		bs, err := output.Type.ByteSize()
		if err != nil {
			return 0, errtypes.WrapErr(errtypes.GenTransErrGetOutputTypeByteSizeFail, err)
		}
		size += bs
	}
	bs, err := output.Lock.ByteSize()
	if err != nil {
		return 0, errtypes.WrapErr(errtypes.GenTransErrGetOutputLockByteSizeFail, err)
	}
	size += bs
	return size, nil
}
