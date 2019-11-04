package types

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

const HexStrPrefix = "0x"

type HexStr struct {
	bytes  []byte
	hexStr string
}

func ParseHexStr(hexStr string) (*HexStr, error) {
	if !strings.HasPrefix(hexStr, HexStrPrefix) {
		return nil, errtypes.WrapErr(errtypes.HexErrNeed0xPrefix, nil)
	}
	if len(hexStr) == 2 {
		return &HexStr{
			bytes:  []byte{},
			hexStr: HexStrPrefix,
		}, nil
	}

	body := hexStr[2:]
	if len(body)%2 == 1 {
		body = "0" + body
	}

	b, err := hex.DecodeString(body)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, err)
	}
	return &HexStr{
		bytes:  b,
		hexStr: HexStrPrefix + body,
	}, nil
}

func NewHexStr(data []byte) *HexStr {
	return &HexStr{
		bytes:  data,
		hexStr: HexStrPrefix + hex.EncodeToString(data),
	}
}

func (hs HexStr) ToHexUint64() (uint64, error) {
	if len(hs.hexStr) > 2 {
		n, err := strconv.ParseUint(hs.hexStr[2:], 16, 64)
		if err != nil {
			return 0, errtypes.WrapErr(errtypes.HexErrStrFormatWrong, err)
		}
		return n, nil
	}
	return 0, errtypes.WrapErr(errtypes.HexErrNeed0xPrefix, nil)
}

func (hs *HexStr) Bytes() []byte {
	return hs.bytes
}

func (hs *HexStr) Hex() string {
	return hs.hexStr
}

func (hs *HexStr) Len() int {
	return len(hs.Bytes())
}

func (hs *HexStr) Append(a HexStr) *HexStr {
	return NewHexStr(append(hs.Bytes(), a.Bytes()...))
}
