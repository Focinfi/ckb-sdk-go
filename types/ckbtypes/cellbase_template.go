package ckbtypes

type CellbaseTemplate struct {
	Cycles interface{} `json:"cycles"`
	Data   Transaction `json:"data"`
	Hash   string      `json:"hash"`
}
