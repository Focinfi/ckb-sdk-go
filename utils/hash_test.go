package utils

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewLockScript(t *testing.T) {
	testPubKey := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	script, err := NewLockScript(testPubKey)
	if err != nil {
		t.Fatal(err)
	}
	expectedScript := &ckbtypes.Script{
		CodeHash: types.BlockAssemblerCodeHash,
		HashType: ckbtypes.HashTypeType,
		Args:     "0xe2fae171d25c36777168caa72dd448677785aa9d",
	}
	assert.Equal(t, expectedScript, script)
}

func TestLockScriptHash(t *testing.T) {
	testPubKey := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	hexStr, err := LockScriptHash(testPubKey)
	if err != nil {
		t.Fatal(err)
	}
	expectedHex := "0x920711df9f85b8cc6638e7fb325c3bee6a19f648d0d74a868feb879d115fe992"
	assert.Equal(t, expectedHex, hexStr.Hex())
}
