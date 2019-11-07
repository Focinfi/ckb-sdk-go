package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func TestNewOutPoint(t *testing.T) {
	type args struct {
		op ckbtypes.OutPoint
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				op: ckbtypes.OutPoint{},
			},
			wantHex: "",
			wantErr: true,
		},
		{
			name: "normal",
			args: args{
				op: ckbtypes.OutPoint{
					Index:  "0x1",
					TxHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
				},
			},
			wantHex: "0x111111111111111111111111111111111111111111111111111111111111111101000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOutPoint(tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOutPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
