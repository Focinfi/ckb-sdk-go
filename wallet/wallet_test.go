package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/rpc"
)

var (
	testPrivKeyHex = "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	testPubKeyHex  = "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	fooPrivKeyHex  = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	fooPubKeyHex   = "0x039166c289b9f905e55f9e3df9f69d7f356b4a22095f894f4715714aa4b56606af"
)

var (
	testCli = rpc.NewClient(rpc.DefaultURL)
	bar, _  = NewWalletByPrivKey(testCli, testPrivKeyHex, true, types.ModeTestNet)
	foo, _  = NewWalletByPrivKey(testCli, fooPrivKeyHex, true, types.ModeTestNet)
)

func TestNewWallet(t *testing.T) {
	w, err := NewWalletByPrivKey(testCli, testPrivKeyHex, true, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, w)
	assert.NotNil(t, w.Key)
	assert.NotNil(t, w.Client)
	assert.NotNil(t, w.lockHashHex)
}

func TestWallet_Balance(t *testing.T) {
	_, err := bar.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestWallet_SendCapacity(t *testing.T) {
	fooAddr, err := foo.Key.Address.Generate()
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("123abc123abc")
	dataHex := types.NewHexStr(data).Hex()
	txHash, err := bar.SendCapacity(context.Background(), fooAddr, 200*types.OneCKBShannon, data, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("transaction hash:", txHash)
		barBalance, fooBalance := balanceOfBarAndFoo()
		t.Logf("balance: bar=%d, foo=%d", barBalance, fooBalance)
		txInfo, err := bar.GetTransaction(context.Background(), txHash)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, dataHex, txInfo.Transaction.OutputsData[0])
	}
}

func TestWallet_DepositToDAO(t *testing.T) {
	outPoint, err := bar.DepositToDAO(context.Background(), 1000*types.OneCKBShannon, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit out point:", outPoint)
}

func balanceOfBarAndFoo() (uint64, uint64) {
	barBalance, _ := bar.Balance(context.Background())
	fooBalance, _ := foo.Balance(context.Background())
	return barBalance, fooBalance
}
