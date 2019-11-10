package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Focinfi/ckb-sdk-go/types"

	"github.com/Focinfi/ckb-sdk-go/rpc"
)

var (
	barPrivKeyHex = "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	barPubKeyHex  = "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	fooPrivKeyHex = "0x58ceea25f67a6baa2c676493fb376347cad88d4208799fb537f31647a8539550"
	fooPubKeyHex  = "0x033f452d7ca46844cd8576bb04ec1e51e2a8cf129da7319435e03fcecbb8bd251e"
)

var (
	testCli    = rpc.NewClient(rpc.DefaultURL)
	bar, _     = NewWalletByPrivKey(testCli, barPrivKeyHex, true, types.ModeTestNet)
	barAddr, _ = bar.Key.Address.Generate()
	foo, _     = NewWalletByPrivKey(testCli, fooPrivKeyHex, true, types.ModeTestNet)
	fooAddr, _ = foo.Key.Address.Generate()
)

func TestNewWallet(t *testing.T) {
	w, err := NewWalletByPrivKey(testCli, barPrivKeyHex, true, types.ModeTestNet)
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

func TestWallet_SendCapacity_MultiSignAddr(t *testing.T) {
	txHash, err := bar.SendCapacity(context.Background(), multiSignWalletAddr, 900*types.OneCKBShannon, nil, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(txHash)
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
