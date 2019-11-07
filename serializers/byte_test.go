package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestByte_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		b       Byte
		wantHex string
	}{
		{
			name:    "0x0",
			b:       Byte(0),
			wantHex: "0x00",
		},
		{
			name:    "0xff",
			b:       Byte(255),
			wantHex: "0xff",
		},
		{
			name:    "0xa",
			b:       Byte(10),
			wantHex: "0x0a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHex := types.NewHexStr(tt.b.Serialize()).Hex()
			assert.Equal(t, tt.wantHex, gotHex)
		})
	}
}
