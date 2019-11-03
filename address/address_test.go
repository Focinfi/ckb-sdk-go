package address

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/addrtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressFromPubKey_Generate(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	expectedShortAddr := "ckt1qyqw97hpw8f9cdnhw95v4fed63yxwau942wsnd3t0u"
	addr, err := NewAddressFromPubKey(testPubKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	actualShortAddr, err := addr.Generate()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedShortAddr, actualShortAddr)
}

func TestAddress_GenerateFullPayloadAddress(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	addr, err := NewAddressFromPubKey(testPubKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		formatType addrtypes.FormatType
		codeHash   string
		args       string
		mode       types.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test net in code format type",
			args: args{
				formatType: addrtypes.FormatTypeCode,
				codeHash:   "0xabcd",
				args:       "0x1234",
				mode:       types.ModeTestNet,
			},
			want:    "ckt1qj4u6y350aq0nk",
			wantErr: false,
		},
		{
			name: "test net in data format type",
			args: args{
				formatType: addrtypes.FormatTypeData,
				codeHash:   "0xabcd",
				args:       "0x1234",
				mode:       types.ModeTestNet,
			},
			want:    "ckt1q24u6y35j0m4ms",
			wantErr: false,
		},
		{
			name: "test net in unsupported format type",
			args: args{
				formatType: addrtypes.FormatType(255),
				codeHash:   "0xabcd",
				args:       "0x1234",
				mode:       types.ModeTestNet,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "test net with wrong hex format code hash",
			args: args{
				formatType: addrtypes.FormatTypeCode,
				codeHash:   "0xabcx",
				args:       "0x1234",
				mode:       types.ModeTestNet,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "test net with wrong hex format args",
			args: args{
				formatType: addrtypes.FormatTypeCode,
				codeHash:   "0xabcx",
				args:       "0x1234",
				mode:       types.ModeTestNet,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addr.GenerateFullPayloadAddress(tt.args.formatType, tt.args.codeHash, tt.args.args, tt.args.mode)
			assert.Equal(t, err != nil, tt.wantErr, "wantErr")
			assert.Equal(t, got, tt.want)
		})
	}
}
