package cellcollector

import (
	"context"
	"testing"

	"github.com/Focinfi/ckb-sdk-go/rpc"
	"github.com/Focinfi/ckb-sdk-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestCellCollector_GetUnspentCells(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	collector := NewCellCollector(rpc.NewClient(rpc.DefaultURL), true)
	lockHash, err := utils.LockScriptHash(testPubKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	cells, totalCap, err := collector.GetUnspentCells(context.Background(), lockHash.Hex(), 1000)
	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, totalCap, uint64(1000))
	t.Logf("totalCap: %d, cells: %v", totalCap, cells)
}

func TestCellCollector_GetUnspentCellsByLockHashes(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	collector := NewCellCollector(rpc.NewClient(rpc.DefaultURL), true)
	lockHash, err := utils.LockScriptHash(testPubKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	cells, totalCap, err := collector.GetUnspentCellsByLockHashes(context.Background(), []string{lockHash.Hex()}, 1000)
	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, totalCap, uint64(1000))
	t.Logf("totalCap: %d, cells: %v", totalCap, cells)
}

func TestCellCollector_GatherInputs(t *testing.T) {
	testPubKeyHex := "0x03579b698bde7d204bdbf845704d0912a56589f61f43d6143d770945c6af350d4e"
	collector := NewCellCollector(rpc.NewClient(rpc.DefaultURL), true)
	lockHash, err := utils.LockScriptHash(testPubKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	inputs, totalCap, err := collector.GatherInputs(context.Background(), []string{lockHash.Hex()}, 100000, 1000, 10, 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, totalCap, uint64(1000))
	t.Logf("totalCap: %d, inputs: %v", totalCap, inputs)
}
