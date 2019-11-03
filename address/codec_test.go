package address

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestEncodeAddress(t *testing.T) {
	hexStr, err := types.ParseHexStr("0xabcd")
	if err != nil {
		t.Fatal(err)
	}
	addr, err := encodeAddress("ckt", hexStr.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ckt140xszvd20l", addr)
}

func TestDecodeAddress(t *testing.T) {
	prefix, payload, err := decodeAddress("ckt140xszvd20l")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ckt", prefix)
	assert.Equal(t, types.NewHexStr(payload).Hex(), "0xabcd")
}
