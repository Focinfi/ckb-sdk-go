package ckbtypes

type Header struct {
	Dao              string `json:"dao"`
	Difficulty       string `json:"difficulty"`
	Epoch            string `json:"epoch"`
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	Number           string `json:"number"`
	ParentHash       string `json:"parent_hash"`
	ProposalsHash    string `json:"proposals_hash"`
	Timestamp        string `json:"timestamp"`
	TransactionsRoot string `json:"transactions_root"`
	UnclesCount      string `json:"uncles_count"`
	UnclesHash       string `json:"uncles_hash"`
	Version          string `json:"version"`
	WitnessesRoot    string `json:"witnesses_root"`
}
