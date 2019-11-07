package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewHexDynVec(t *testing.T) {
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
			name: "empty slice",
			args: args{
				hexes: []string{},
			},
			wantHex: "0x04000000",
			wantErr: false,
		},
		{
			name: "[0xabc,0x1]",
			args: args{
				hexes: []string{"0xabc", "0x1"},
			},
			wantHex: "0x170000000c00000012000000020000000abc0100000001",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHexDynVec(tt.args.hexes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHexDynVec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotStr := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotStr)
			}
		})
	}
}

func TestNewByteDynVecByHexes(t *testing.T) {
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
			name:    "empty slice",
			args:    args{hexes: []string{}},
			wantHex: "0x04000000",
			wantErr: false,
		},
		{
			name:    "invalid hex",
			args:    args{hexes: []string{"0xxx"}},
			wantHex: "",
			wantErr: true,
		},
		{
			name:    "normal",
			args:    args{hexes: []string{"0x123", "0xab"}},
			wantHex: "0x170000000c0000001200000002000000012301000000ab",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByteDynVecByHexes(tt.args.hexes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByteDynVecByHexes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
