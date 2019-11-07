package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewWitnessesByHexes(t *testing.T) {
	type args struct {
		hexes []string
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name: "[]",
			args: args{
				hexes: nil,
			},
			wantHex: "0x04000000",
			wantErr: false,
		},
		{
			name: "[0x]",
			args: args{
				hexes: []string{"0x"},
			},
			wantHex: "0x0c0000000800000000000000",
			wantErr: false,
		},
		{
			name: "[0x01,0xabc]",
			args: args{
				hexes: []string{"0x01", "0xabc"},
			},
			wantHex: "0x170000000c000000110000000100000001020000000abc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWitnessesByHexes(tt.args.hexes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWitnessesByHexes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantHex, types.NewHexStr(got.Serialize()).Hex())
		})
	}
}
