package ckbtypes

type CellInfo struct {
	Cell   LiveCell `json:"cell"`
	Status string   `json:"status"`
}
