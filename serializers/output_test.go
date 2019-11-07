package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewOutputsDynVec(t *testing.T) {
	type args struct {
		outputs []ckbtypes.Output
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
				outputs: []ckbtypes.Output{},
			},
			wantHex: "0x04000000",
			wantErr: false,
		},
		{
			name: "nil type script",
			args: args{
				outputs: []ckbtypes.Output{
					{
						Capacity: "0x1",
						Lock: ckbtypes.Script{
							Args:     "0xabcd",
							CodeHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
							HashType: ckbtypes.HashTypeType,
						},
					},
				},
			},
			wantHex: "0x57000000080000004f00000010000000180000004f00000001000000000000003700000010000000300000003100000011111111111111111111111111111111111111111111111111111111111111110102000000abcd",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOutputsDynVec(tt.args.outputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOutputsDynVec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
