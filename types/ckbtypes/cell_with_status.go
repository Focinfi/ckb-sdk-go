package ckbtypes

type CellWithStatus struct {
	Cell   CellInfo `json:"cell"`
	Status string   `json:"status"`
}
