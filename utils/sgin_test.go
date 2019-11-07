package utils

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/key"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
)

func TestSignTransaction(t *testing.T) {
	testPrivKeyHex := "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	k, err := key.NewFromPrivKeyHex(testPrivKeyHex, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		key         key.Key
		transaction *ckbtypes.Transaction
	}
	tests := []struct {
		name        string
		args        args
		wantWitness []string
		wantErr     bool
	}{
		{
			name: "normal",
			args: args{
				key: *k,
				transaction: &ckbtypes.Transaction{
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
					Witnesses:   EmptyWitnessesByLen(1),
					OutputsData: []string{"0x", "0x"},
					Version:     "0x0",
				},
			},
			wantWitness: []string{"0x739942d6b2214a549ddcb23f135df85cf3f1b4297f1c27e0ceaa56bee88683a51de83016a61f65d8ea9bafa822d5a6f7b86db5d460693acc846a8dbf525a708301"},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SignTransaction(tt.args.key, tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("SignTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, tt.wantWitness, tt.args.transaction.Witnesses)
			}
		})
	}
}
