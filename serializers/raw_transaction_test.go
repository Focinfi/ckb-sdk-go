package serializers

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewRawTransaction(t *testing.T) {
	type args struct {
		transaction ckbtypes.Transaction
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
				transaction: ckbtypes.Transaction{
					CellDeps: []ckbtypes.CellDep{
						{
							DepType: ckbtypes.DepTypeDepGroup,
							OutPoint: ckbtypes.OutPoint{
								TxHash: "0xb815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c70",
								Index:  "0x0",
							},
						},
					},
					HeaderDeps: []string{},
					Inputs: []ckbtypes.Input{
						{
							PreviousOutput: ckbtypes.OutPoint{
								Index:  "0x0",
								TxHash: "0x95202d82c38ad9544f09df04b4e5a161038248376f4143fb2856ad2d59b11a68",
							},
							Since: "0x0",
						},
					},
					Outputs: []ckbtypes.Output{
						{
							Capacity: "0x4b4038a00",
							Lock: ckbtypes.Script{
								Args:     "0xddbd7f09eb480450c1b1ed2c8696248de91c6802",
								CodeHash: "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
								HashType: ckbtypes.HashTypeType,
							},
						},
						{
							Capacity: "0x2ab7471bdf",
							Lock: ckbtypes.Script{
								Args:     "0xe2fae171d25c36777168caa72dd448677785aa9d",
								CodeHash: "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
								HashType: ckbtypes.HashTypeType,
							},
						},
					},
					OutputsData: []string{"0x", "0x"},
					Version:     "0x0",
				},
			},
			wantHex: "0x5f0100001c00000020000000490000004d0000007d0000004b0100000000000001000000b815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c7000000000010000000001000000000000000000000095202d82c38ad9544f09df04b4e5a161038248376f4143fb2856ad2d59b11a6800000000ce0000000c0000006d00000061000000100000001800000061000000008a03b404000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000ddbd7f09eb480450c1b1ed2c8696248de91c680261000000100000001800000061000000df1b47b72a000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000e2fae171d25c36777168caa72dd448677785aa9d140000000c000000100000000000000000000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRawTransaction(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRawTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex := types.NewHexStr(got.Serialize()).Hex()
				assert.Equal(t, tt.wantHex, gotHex)
			}
		})
	}
}
