package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func TestNewWitnessArgs(t *testing.T) {
	type args struct {
		witness ckbtypes.Witness
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name:    "empty",
			wantHex: "0x10000000100000001000000010000000",
			wantErr: false,
		},
		{
			name: "normal",
			args: args{
				witness: ckbtypes.Witness{
					Lock:       "0x123",
					InputType:  "0xabc",
					OutputType: "0x321",
				},
			},
			wantHex: "0x2200000010000000160000001c000000020000000123020000000abc020000000321",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWitnessArgs(tt.args.witness)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWitnessArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotHex := types.NewHexStr(got.Serialize()).Hex()
			assert.Equal(t, tt.wantHex, gotHex)
		})
	}
}
