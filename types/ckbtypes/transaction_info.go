package ckbtypes

type TransactionInfo struct {
	Transaction Transaction           `json:"transaction"`
	Status      TransactionStatusInfo `json:"tx_status"`
}

type TransactionStatusInfo struct {
	BlockHash string            `json:"block_hash"`
	Status    TransactionStatus `json:"status"`
}
