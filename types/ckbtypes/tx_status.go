package ckbtypes

type TxStatus struct {
	BlockHash string `json:"block_hash"`
	Status    string `json:"status"`
}
