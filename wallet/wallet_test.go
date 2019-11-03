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
)

func TestNewWallet(t *testing.T) {
	client := rpc.NewClient(rpc.DefaultURL)
	wallet, err := NewWalletByPrivKey(client, testPrivKeyHex, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wallet:", wallet)

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
}
