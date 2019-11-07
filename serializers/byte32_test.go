package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestByte32(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name:    "empty string",
			args:    args{hex: ""},
			wantHex: "0x",
			wantErr: true,
		},
		{
			name:    "0x",
			args:    args{hex: "0x"},
			wantHex: "0x",
			wantErr: true,
		},
		{
			name:    "length less than 32 bytes",
			args:    args{hex: "0xabc"},
			wantHex: "0x",
			wantErr: true,
		},
		{
			name:    "normal",
			args:    args{hex: "0x1111111111111111111111111111111111111111111111111111111111111111"},
			wantHex: "0x1111111111111111111111111111111111111111111111111111111111111111",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByte32(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByte32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotHex := types.NewHexStr(got.Serialize()).Hex()
			assert.Equal(t, tt.wantHex, gotHex)
		})
	}
}
