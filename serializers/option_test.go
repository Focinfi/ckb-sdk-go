package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewOptionByteFixVec(t *testing.T) {
	type args struct {
		hexStr string
	}
	tests := []struct {
		name       string
		args       args
		wantVecHex string
		wantErr    bool
	}{
		{
			name:       "empty hex string",
			wantVecHex: "0x",
			wantErr:    false,
		},
		{
			name:       "not a hex string",
			args:       args{hexStr: "xxx"},
			wantVecHex: "",
			wantErr:    true,
		},
		{
			name:       "hex string",
			args:       args{hexStr: "0xabc"},
			wantVecHex: "0x020000000abc",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVec, err := NewByteFixVecOption(tt.args.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByteFixVecOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(gotVec.Serialize()).Hex()
				assert.Equal(t, tt.wantVecHex, gotHex)
			}
		})
	}
}
