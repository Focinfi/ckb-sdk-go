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
		transaction ckbtypes.Transaction
	}
	tests := []struct {
		name        string
		args        args
		wantWitness []interface{}
		wantErr     bool
	}{
		{
			name: "normal",
			args: args{
				key: *k,
				transaction: ckbtypes.Transaction{
					CellDeps: []ckbtypes.CellDep{
						{
							DepType: ckbtypes.DepTypeDepGroup,
							OutPoint: ckbtypes.OutPoint{
								TxHash: "0xb815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c70",
								Index:  "0x0",
							},
						},
						{
							DepType: ckbtypes.DepTypeCode,
							OutPoint: ckbtypes.OutPoint{
								TxHash: "0xb5724acb4f5f82afb717c3ec3fe025d3b6e45ff48f4ffbb6162c950399cbcabe",
								Index:  "0x2",
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
						{
							PreviousOutput: ckbtypes.OutPoint{
								Index:  "0x0",
								TxHash: "0x95202d82c38ad9544f09df04b4e5a161038248376f4143fb2856ad2d59b11a69",
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
							Type: &ckbtypes.Script{
								Args:     "0x",
								CodeHash: "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e",
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
					Witnesses:   []interface{}{ckbtypes.Witness{}, "0x"},
					OutputsData: []string{"0x0000000000000000", "0x"},
					Version:     "0x0",
				},
			},
			wantWitness: []interface{}{
				"0x5500000010000000550000005500000041000000d48be96d5c2281aa2dd04f5242467043220c5847b691b033da7fdf1c41962c724abff0be6d83507bdefa9a4f2f1c7fc513dd7f58ba97894913cc63ff274c03e100",
				"0x",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := SignTransaction(tt.args.key, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, tt.wantWitness, tx.Witnesses)
			}
		})
	}
}
