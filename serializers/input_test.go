package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func TestNewInputsFixVec(t *testing.T) {
	type args struct {
		inputs []ckbtypes.Input
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name: "empty inputs",
			args: args{
				inputs: []ckbtypes.Input{},
			},
			wantHex: "0x00000000",
			wantErr: false,
		},
		{
			name: "normal",
			args: args{
				inputs: []ckbtypes.Input{
					{
						PreviousOutput: ckbtypes.OutPoint{
							Index:  "0x1",
							TxHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
						},
						Since: "0x01",
					},
				},
			},
			wantHex: "0x010000000100000000000000111111111111111111111111111111111111111111111111111111111111111101000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInputsFixVec(tt.args.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInputsFixVec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
