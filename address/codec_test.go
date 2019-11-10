package address

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestEncodeAddress(t *testing.T) {
	hexStr, err := types.ParseHexStr("0x0101226cc6c280b694f9959443e46f715fcc7c148156")
	if err != nil {
		t.Fatal(err)
	}
	addr, err := EncodeAddress("ckt", hexStr.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ckt1qyqjymxxc2qtd98ejk2y8er0w90uclq5s9tqzyzs0j", addr)
}

func TestDecodeAddress(t *testing.T) {
	prefix, payload, err := DecodeAddress("ckt140xszvd20l")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ckt", prefix)
	assert.Equal(t, types.NewHexStr(payload).Hex(), "0xabcd")
}
