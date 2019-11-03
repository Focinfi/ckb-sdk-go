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
	assert.Equal(t, expectedDigestedHex, types.NewHexStr(d).String())
}
