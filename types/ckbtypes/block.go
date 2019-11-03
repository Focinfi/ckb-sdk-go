package ckbtypes

type Block struct {
	Header       Header        `json:"header"`
	Proposals    []string      `json:"proposals"`
	Transactions []Transaction `json:"transactions"`
	Uncles       []UncleBlock  `json:"uncles"`
}
