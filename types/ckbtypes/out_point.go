package ckbtypes

type OutPoint struct {
	Index  string `json:"index"`
	TxHash string `json:"tx_hash"`
}
