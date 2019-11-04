package key

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/stretchr/testify/assert"
)

var (
	testPrivKeyHex = "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	testPubKeyHex  = "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
)

func TestNewFromPrivKeyHex(t *testing.T) {
	key, err := NewFromPrivKeyHex(testPrivKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testPubKeyHex, key.pubKeyHex.Hex())
}

func TestSign(t *testing.T) {
	key, err := NewFromPrivKeyHex(testPrivKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	gotSign, err := key.Sign("0x1abc")
	if err != nil {
		t.Fatal(err)
	}
	wantSign := "0x30450221009e0f9776b125da2ffd17aaef11aaf2c68779abf0df79b7c3b11c22b6dcca9c8b02201cff47a3adc87f4e02c17da0a1b0540e2e08a7c50d0a93568cef33583e75b1fb"
	assert.Equal(t, wantSign, gotSign.Hex())
}

func TestSignRecoverable(t *testing.T) {
	key, err := NewFromPrivKeyHex(testPrivKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	gotSign, err := key.SignRecoverableFor32BytesHex("0x1000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	wantSign := "0xdf14f66e8f3ab10fedee7147c8ce3be9044da14fd8e4cd0dbd654e7bf59e3ce42ce9fa8e5b89af24bb1a8855ded3ad99f05fd35a9d07a4225b89a39be127aca401"
	assert.Equal(t, wantSign, gotSign.Hex())
}
