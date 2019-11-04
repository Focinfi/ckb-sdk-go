package wallet

import (
	"context"
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/rpc"
)

var (
	testPrivKeyHex = "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	testPubKeyHex  = "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	fooPrivKeyHex  = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	fooPubKeyHex   = "0x039166c289b9f905e55f9e3df9f69d7f356b4a22095f894f4715714aa4b56606af"
)

func TestNewWallet(t *testing.T) {
	client := rpc.NewClient(rpc.DefaultURL)
	wallet, err := NewWalletByPrivKey(client, testPrivKeyHex, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wallet:", wallet)
	foo, err := NewWalletByPrivKey(client, fooPrivKeyHex, true, types.ModeTestNet)

	balance, err := wallet.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("balance:", balance)

	unspentCells, err := wallet.GetUnspentCells(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("unspentCells:", unspentCells)

	fooAddr, err := foo.Key.Address.Generate()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := wallet.GenerateTx(context.Background(), fooAddr, 1000*types.OneCKBShannon, nil, types.OneCKBShannon, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("tx:", tx)
}
