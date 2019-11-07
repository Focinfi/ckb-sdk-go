package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewScript(t *testing.T) {
	type args struct {
		script ckbtypes.Script
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				script: ckbtypes.Script{
					Args:     "0xabcd",
					CodeHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
					HashType: ckbtypes.HashTypeType,
				},
			},
			wantHex: "0x3700000010000000300000003100000011111111111111111111111111111111111111111111111111111111111111110102000000abcd",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewScript(tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
