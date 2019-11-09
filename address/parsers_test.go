package address

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/stretchr/testify/assert"
)

func TestParseShortPayloadAddress(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	expectedShortAddr := "ckt1qyqw97hpw8f9cdnhw95v4fed63yxwau942wsnd3t0u"
	addr, err := NewAddressFromPubKey(testPubKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	data, err := ParseShortPayloadAddress(expectedShortAddr, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := AddrConfig{
		FormatType:    addrtypes.FormatTypeShortLock,
		CodeHashIndex: addrtypes.CodeHashIndex0,
		Args:          addr.PubKey.Blake160,
	}
	assert.Equal(t, expectedData, data)
}

func TestParseFullPayloadAddress(t *testing.T) {
	type args struct {
		address string
		mode    types.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		//{
		//	name: ""
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFullPayloadAddress(tt.args.address, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFullPayloadAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
