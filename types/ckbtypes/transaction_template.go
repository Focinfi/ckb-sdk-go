package ckbtypes

type TransactionTemplate struct {
	Hash     string      `json:"hash"`
	Required bool        `json:"required"`
	Cycles   string      `json:"cycles"`
	Depends  []string    `json:"depends"`
	Data     Transaction `json:"data"`
}
