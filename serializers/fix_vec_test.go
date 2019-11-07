package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewByteFixVecByHex(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name:    "empty string",
			args:    args{str: ""},
			wantHex: "",
			wantErr: true,
		},
		{
			name:    "not a hex string",
			args:    args{str: "0a123"},
			wantHex: "",
			wantErr: true,
		},
		{
			name:    "normal",
			args:    args{str: "0xabc"},
			wantHex: "0x020000000abc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByteFixVecByHex(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByteFixVecByHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}

func TestNewByte32FixVecByHexes(t *testing.T) {
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
			wantHex: "0x00000000",
			wantErr: false,
		},
		{
			name:    "validate byte32 items",
			args:    args{hexes: []string{"0xabc"}},
			wantHex: "",
			wantErr: true,
		},
		{
			name: "normal",
			args: args{
				hexes: []string{
					"0x1111111111111111111111111111111111111111111111111111111111111111",
					"0x2222222222222222222222222222222222222222222222222222222222222222",
				}},
			wantHex: "0x0200000011111111111111111111111111111111111111111111111111111111111111112222222222222222222222222222222222222222222222222222222222222222",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByte32FixVecByHexes(tt.args.hexes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByte32FixVecByHexes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
