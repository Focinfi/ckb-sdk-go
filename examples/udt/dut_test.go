package udt

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/Focinfi/ckb-sdk-go/serializers"

	"github.com/Focinfi/ckb-sdk-go/crypto/blake2b"

	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/types"
	"github.com/Focinfi/ckb-sdk-go/types/ckbtypes"
	"github.com/Focinfi/ckb-sdk-go/wallet"
)

var (
	barPrivKeyHex = "0x3f86634c419dd7f266793c9fda9fb4ccbe121ce395ed14e699a741a4dabf0177"
	barPubKeyHex  = "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	fooPrivKeyHex = "0x58ceea25f67a6baa2c676493fb376347cad88d4208799fb537f31647a8539550"
	fooPubKeyHex  = "0x033f452d7ca46844cd8576bb04ec1e51e2a8cf129da7319435e03fcecbb8bd251e"

	duktapeTxHash = "0xa44c330700799190240168aa050b81ad3ecba0e4a00c3cf7894c32aad1c7a4b1"
)

var (
	testCli    = rpc.NewClient(rpc.DefaultURL)
	bar, _     = wallet.NewWalletByPrivKey(testCli, barPrivKeyHex, true, ckbtypes.HashTypeType, types.ModeTestNet)
	barAddr, _ = bar.Key.Address.Generate()
	foo, _     = wallet.NewWalletByPrivKey(testCli, fooPrivKeyHex, true, ckbtypes.HashTypeType, types.ModeTestNet)
	fooAddr, _ = foo.Key.Address.Generate()
)

func TestUDT_Init(t *testing.T) {
	// duktape data
	duktapData, err := ioutil.ReadFile("./duktape")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("duktapData size:", len(duktapData))

	//send the duktapedata into chain
	hashTx, err := bar.SendCapacity(context.Background(), barAddr, 300000*types.OneCKBShannon, duktapData, types.OneCKBShannon)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("duktapData transaction hash:", hashTx)
}

func TestUDT_Deploy(t *testing.T) {
	// duktape data
	duktapData, err := ioutil.ReadFile("./duktape")
	if err != nil {
		t.Fatal(err)
	}
	duktapeCodeHash, err := blake2b.Digest(duktapData)
	if err != nil {
		t.Fatal(err)
	}
	duktapeCodeHashHex := types.NewHexStr(duktapeCodeHash)

	duktapeCellDep := ckbtypes.CellDep{
		DepType: ckbtypes.DepTypeCode,
		OutPoint: ckbtypes.OutPoint{
			Index:  types.Hex0.Hex(),
			TxHash: duktapeTxHash,
		},
	}

	genesisUDTCount := types.NewHexStr(types.HexUint64(1000000).LittleEndianBytes(4))
	genesisTx, err := bar.GenerateTx(context.Background(), barAddr, 20000*types.OneCKBShannon, genesisUDTCount.Bytes(), types.OneCKBShannon, true)
	if err != nil {
		t.Fatal(err)
	}
	firstInputBytes, err := serializers.NewInput(genesisTx.Inputs[0])
	if err != nil {
		t.Fatal(err)
	}
	firstInputHex := types.NewHexStr(firstInputBytes.Serialize())
	code, err := ioutil.ReadFile("./demo.js")
	if err != nil {
		t.Fatal(err)
	}
	codeHex := types.NewHexStr(code)
	duktapeUDTScript := ckbtypes.Script{
		Args:     codeHex.Hex() + firstInputHex.Hex(),
		CodeHash: duktapeCodeHashHex.Hex(),
		HashType: ckbtypes.HashTypeData,
	}

	t.Log(duktapeCellDep)

	slr, err := serializers.NewScript(duktapeUDTScript)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(slr.Serialize()))
}
