package ckbtypes

type BlockHeader struct {
	CompactTarget    string `json:"compact_target"`
	Hash             string `json:"hash"`
	Number           string `json:"number"`
	ParentHash       string `json:"parent_hash"`
	Nonce            string `json:"nonce"`
	Timestamp        string `json:"timestamp"`
	TransactionsRoot string `json:"transactions_root"`
	ProposalsHash    string `json:"proposals_hash"`
	UnclesHash       string `json:"uncles_hash"`
	Version          string `json:"version"`
	Epoch            string `json:"epoch"`
	DAO              string `json:"dao"`
}
