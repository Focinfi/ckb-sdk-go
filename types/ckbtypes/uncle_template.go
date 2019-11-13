package ckbtypes

type UncleTemplate struct {
	Hash      string      `json:"hash"`
	Required  bool        `json:"required"`
	Proposals []string    `json:"proposals"`
	Header    BlockHeader `json:"header"`
}
