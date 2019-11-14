package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHexStr(t *testing.T) {
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
			name:    "all 0",
			args:    args{hexStr: "0x00000000"},
			wantHex: "0x00000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHexStr(tt.args.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.Equal(t, tt.wantHex, got.Hex())
			}
		})
	}
}
