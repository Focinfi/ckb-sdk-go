package ckbtypes

type UncleBlock struct {
	CellBase  Transaction `json:"cellbase"`
	Header    Header      `json:"header"`
	Proposals []string    `json:"proposals"`
}
