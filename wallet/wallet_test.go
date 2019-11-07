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
	minner, err := NewWalletByPrivKey(client, testPrivKeyHex, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("minner:", minner)
	foo, err := NewWalletByPrivKey(client, fooPrivKeyHex, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := minner.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("miner balance:", balance)
	fooBalance, err := foo.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("foo balance:", fooBalance)

	unspentCells, err := minner.GetUnspentCells(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("unspentCells:", unspentCells)

	fooAddr, err := foo.Key.Address.Generate()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := minner.GenerateTx(context.Background(), fooAddr, 1000*types.OneCKBShannon, nil, types.OneCKBShannon, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("tx:", tx)

	txHash, err := minner.SendCapacity(context.Background(), fooAddr, 1000*types.OneCKBShannon, nil, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transaction hash:", txHash)
}
