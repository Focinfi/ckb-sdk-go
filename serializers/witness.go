package serializers

import "github.com/Focinfi/ckb-sdk-go/types/errtypes"

type Witness = ByteFixVec

type Witnesses = ByteDynVec

func NewWitnessesByHexes(hexes []string) (*ByteDynVec, error) {
	return NewByteDynVecByHexes(hexes)
}

func NewWitnessesByInterfaces(interfaces []interface{}) (*ByteDynVec, error) {
	hexStr := make([]string, 0, len(interfaces))
	for _, hex := range interfaces {
		switch h := hex.(type) {
		case string:
			hexStr = append(hexStr, h)
		default:
			return nil, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, nil)
		}
	}

	return NewByteDynVecByHexes(hexStr)
}
