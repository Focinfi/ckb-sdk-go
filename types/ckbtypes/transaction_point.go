package ckbtypes

type TransactionPoint struct {
	BlockNumber string `json:"block_number"`
	Index       string `json:"index"`
	TxHash      string `json:"tx_hash"`
}
