package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewUint32ByHex(t *testing.T) {
	type args struct {
		hexStr string
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name:    "0x0",
			args:    args{hexStr: "0x0"},
			wantHex: "0x00000000",
			wantErr: false,
		},
		{
			name:    "0x1",
			args:    args{hexStr: "0x1"},
			wantHex: "0x01000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUint32ByHex(tt.args.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint32ByHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
