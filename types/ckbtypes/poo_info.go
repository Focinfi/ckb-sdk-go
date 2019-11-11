package ckbtypes

type PoolInfo struct {
	LastTxsUpdatedAt string `json:"last_txs_updated_at"`
	Orphan           string `json:"orphan"`
	Pending          string `json:"pending"`
	Proposed         string `json:"proposed"`
	TotalTxCycles    string `json:"total_tx_cycles"`
	TotalTxSize      string `json:"total_tx_size"`
}
