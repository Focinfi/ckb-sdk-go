package ckbtypes

type Script struct {
	Args     string   `json:"args"`
	CodeHash string   `json:"code_hash"`
	HashType HashType `json:"hash_type"`
}
