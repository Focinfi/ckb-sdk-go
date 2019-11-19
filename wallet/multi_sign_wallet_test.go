package wallet

import (
	"context"
	"testing"

	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

var (
	multiSignWalletConfig = MultiSignWalletConfig{
		RequireN:  2,
		Threshold: 2,
		PubKeys:   []string{barPubKeyHex, fooPubKeyHex},
	}

	multiSignWallet, _     = NewMultiSignWallet(*rpc.NewClient(rpc.DefaultURL), multiSignWalletConfig, true, types.ModeTestNet)
	multiSignWalletAddr, _ = multiSignWallet.Address()
)

func TestMultiSignWalletConfig(t *testing.T) {
	type args struct {
		requireN  uint8
		threshold uint8
		pubKeys   []string
	}
	tests := []struct {
		name            string
		args            args
		wantBlake160Hex string
		wantErr         bool
	}{
		{
			name: "pubKeys number greater than 255",
			args: args{
				requireN:  1,
				threshold: 1,
				pubKeys:   make([]string, 256),
			},
			wantBlake160Hex: "",
			wantErr:         true,
		},
		{
			name: "normal",
			args: args{
				requireN:  1,
				threshold: 2,
				pubKeys:   []string{barPubKeyHex, fooPubKeyHex},
			},
			wantBlake160Hex: "0x226cc6c280b694f9959443e46f715fcc7c148156",
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMultiSignWalletConfig(tt.args.requireN, tt.args.threshold, tt.args.pubKeys, 100)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMultiSignWalletConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				gotHex, err := got.Blake160()
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantBlake160Hex, types.NewHexStr(gotHex).Hex())
			}
		})
	}
}

func TestNewMultiSignWallet(t *testing.T) {
	wallet, err := NewMultiSignWallet(*rpc.NewClient(rpc.DefaultURL), multiSignWalletConfig, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, wallet.sysCells)
	assert.NotNil(t, wallet.lock)
	assert.NotNil(t, wallet.lockHash)
}

func TestMultiSignWallet_Address(t *testing.T) {
	addr, err := multiSignWallet.Address()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ckt1qyqjymxxc2qtd98ejk2y8er0w90uclq5s9tqzyzs0j", addr)
}

func TestMultiSignWallet_GetBalance(t *testing.T) {
	balance, err := multiSignWallet.GetBalance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("balance: %d", balance)
}

func TestMultiSignWallet_SendCapacity(t *testing.T) {
	type args struct {
		targetAddr string
		capacity   uint64
		privKeys   []string
		data       []byte
		fee        uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "not enough privKeys",
			args: args{
				targetAddr: fooAddr,
				capacity:   60 * types.OneCKBShannon,
				privKeys:   []string{barPrivKeyHex},
				fee:        types.OneCKBShannon,
			},
			wantErr: true,
		},
		{
			name: "not signed by bar and foo",
			args: args{
				targetAddr: fooAddr,
				capacity:   100 * types.OneCKBShannon,
				privKeys:   []string{barPrivKeyHex, barPrivKeyHex},
				data:       []byte{5},
				fee:        types.OneCKBShannon,
			},
			wantErr: true,
		},
		{
			name: "signed by foo and bar",
			args: args{
				targetAddr: fooAddr,
				capacity:   110 * types.OneCKBShannon,
				privKeys:   []string{barPrivKeyHex, fooPrivKeyHex},
				data:       []byte{5},
				fee:        types.OneCKBShannon,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet := multiSignWallet
			got, err := wallet.SendCapacity(context.Background(), tt.args.targetAddr, tt.args.capacity, tt.args.privKeys, tt.args.data, tt.args.fee)
			t.Log("txHash:", got, "err:", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiSignWallet.SendCapacity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MultiSignWallet.SendCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
