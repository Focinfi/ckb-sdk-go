package blake160

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestBlake160(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	d, err := Blake160(testPubKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	expectedDigestedHex := "0xe2fae171d25c36777168caa72dd448677785aa9d"
	assert.Equal(t, expectedDigestedHex, types.NewHexStr(d).Hex())
}

func TestBlake160_Sign(t *testing.T) {
	meta := []byte{1, 2, 2}
	bar, _ := Blake160("0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e")
	foo, _ := Blake160("0x039166c289b9f905e55f9e3df9f69d7f356b4a22095f894f4715714aa4b56606af")
	body := append(meta, bar...)
	body = append(body, foo...)
	h, _ := Blake160Binary(body)
	t.Log(types.NewHexStr(h).Hex())
}
