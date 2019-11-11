package ckbtypes

type BlockHashIndexState struct {
	BlockHash   string `json:"block_hash"`
	BlockNumber string `json:"block_number"`
	LockHash    string `json:"lock_hash"`
}
