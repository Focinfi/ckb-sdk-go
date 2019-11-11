package rpc

import (
	"testing"
)

var client = NewClient(DefaultURL)

func TestClient(t *testing.T) {
	//n, err := client.GetTipBlockNumber(context.Background())
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_tip_block_number: %d", n)
	//
	//tipHeader, err := client.GetTipHeader(context.Background())
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_tip_header: %#v", tipHeader)
	//
	//currentEpoch, err := client.GetCurrentEpoch(context.Background())
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_current_epoch: %#v", currentEpoch)

	//epoch, err := client.GetEpochByNumber(context.Background(), 0)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_epoch_by_number: %#v", epoch)

	//hash, err := client.GetBlockHash(context.Background(), 1)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_block_hash: %v", hash)

	//hash, err := client.GetBlock(context.Background(), "0xed707a2e9b367ce685f7dbc378e8849163ef07e5c6b2d98ffafa2df422dd2683")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_block: %#v", hash)

	//hash, err := client.GetBlockByNumber(context.Background(), 1)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_block_by_number: %#v", hash)

	//header, err := client.GetHeader(context.Background(), "0xed707a2e9b367ce685f7dbc378e8849163ef07e5c6b2d98ffafa2df422dd2683")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_header: %#v", header)

	//header, err := client.GetHeaderByNumber(context.Background(), 1)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_header_by_number: %#v", header)

	//header, err := client.GetCellsByLockHash(context.Background(), "0x920711df9f85b8cc6638e7fb325c3bee6a19f648d0d74a868feb879d115fe992", 100, 200)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_cells_by_lock_hash: %#v", header)

	//header, err := client.GetLiveCellsByLockHash(context.Background(), "0x920711df9f85b8cc6638e7fb325c3bee6a19f648d0d74a868feb879d115fe992", 100, 200)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_live_cells_by_lock_hash: %#v", header)

	//liveCell, err := client.GetLiveCell(context.Background(), types.OutPoint{
	//	TxHash: "0x45414d03927d0c4fcbd868f78aa0a5e63ac675a51cdf59999ace2fbcaa48e060",
	//	Index:  "0x0",
	//}, true)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_live_cell: %#v", liveCell)

	//transaction, err := client.GetTransaction(context.Background(), "0x45414d03927d0c4fcbd868f78aa0a5e63ac675a51cdf59999ace2fbcaa48e060")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("get_transaction: %#v", transaction)
}
