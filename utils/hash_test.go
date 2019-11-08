package utils

import (
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/stretchr/testify/assert"
)

func TestNewLockScript(t *testing.T) {
	testPubKey := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	script, err := NewLockScript(testPubKey)
	if err != nil {
		t.Fatal(err)
	}
	expectedScript := &ckbtypes.Script{
		CodeHash: types.BlockAssemblerCodeHash,
		HashType: ckbtypes.HashTypeType,
		Args:     "0xe2fae171d25c36777168caa72dd448677785aa9d",
	}
	assert.Equal(t, expectedScript, script)
}

func TestLockScriptHash(t *testing.T) {
	testPubKey := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	hexStr, err := LockScriptHash(testPubKey)
	if err != nil {
		t.Fatal(err)
	}
	expectedHex := "0x920711df9f85b8cc6638e7fb325c3bee6a19f648d0d74a868feb879d115fe992"
	assert.Equal(t, expectedHex, hexStr.Hex())
}

func TestRawTransactionHash(t *testing.T) {
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
					Witnesses:   []interface{}{},
					OutputsData: []string{"0x", "0x"},
					Version:     "0x0",
				},
			},
			wantHex: "0xd54cc789afb3b7153ce9032a9227341011f1287b802fc56d8b0f428b66a3c503",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawTransactionHash(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawTransactionHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.Equal(t, tt.wantHex, got.Hex())
			}
		})
	}
}
