package wallet

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"

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
	bar, _     = NewWalletByPrivKey(testCli, barPrivKeyHex, true, ckbtypes.HashTypeType, types.ModeTestNet)
	barAddr, _ = bar.Key.Address.Generate()
	foo, _     = NewWalletByPrivKey(testCli, fooPrivKeyHex, true, ckbtypes.HashTypeType, types.ModeTestNet)
	fooAddr, _ = foo.Key.Address.Generate()
)

func TestNewWallet(t *testing.T) {
	w, err := NewWalletByPrivKey(testCli, barPrivKeyHex, true, ckbtypes.HashTypeType, types.ModeTestNet)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, w)
	assert.NotNil(t, w.Key)
	assert.NotNil(t, w.Client)
	assert.NotNil(t, w.lockHashHex)
}

func TestWallet_Balance(t *testing.T) {
	balance, err := bar.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("balance:", balance)
}

func TestWallet_SendCapacity(t *testing.T) {
	data := []byte("123abc123abc")
	dataHex := types.NewHexStr(data).Hex()
	txHash, err := bar.SendCapacity(context.Background(), barAddr, 200*types.OneCKBShannon, data, types.OneCKBShannon)
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
	t.Log(outPoint)
}

func TestWallet_StartWithdrawingFromDAO(t *testing.T) {
	depositOutPoint := ckbtypes.OutPoint{
		Index:  "0x0",
		TxHash: "0x89dbfcd76f4d1ab592e97c99b86f93160b8f403c7228f639f66ff8ab8d44c3a3",
	}
	withdrawOutPoint, err := bar.StartWithdrawingFromDAO(context.Background(), depositOutPoint, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(withdrawOutPoint)
}

// 0x657a6d0cc99ee3ed015d48301f0b122f200a429dc1e75350b0f91f4bdfe93935
func TestWallet_GenWithdrawFromDAOTx(t *testing.T) {
	depositOutPoint := ckbtypes.OutPoint{
		Index:  "0x0",
		TxHash: "0x89dbfcd76f4d1ab592e97c99b86f93160b8f403c7228f639f66ff8ab8d44c3a3",
	}
	withdrawOutPoint := ckbtypes.OutPoint{
		Index:  "0x0",
		TxHash: "0x657a6d0cc99ee3ed015d48301f0b122f200a429dc1e75350b0f91f4bdfe93935",
	}
	tx, err := bar.GenWithdrawFromDAOTx(context.Background(), depositOutPoint, withdrawOutPoint, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	txJSON, _ := json.MarshalIndent(tx, "", "  ")
	t.Log(string(txJSON))
}

func balanceOfBarAndFoo() (uint64, uint64) {
	barBalance, _ := bar.Balance(context.Background())
	fooBalance, _ := foo.Balance(context.Background())
	return barBalance, fooBalance
}
