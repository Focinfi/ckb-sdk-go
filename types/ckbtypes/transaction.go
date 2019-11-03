package ckbtypes

type Transaction struct {
	CellDeps    []CellDep `json:"cell_deps"`
	Hash        string    `json:"hash"`
	HeaderDeps  []string  `json:"header_deps"`
	Inputs      []Input   `json:"inputs"`
	Outputs     []Output  `json:"outputs"`
	OutputsData []string  `json:"outputs_data"`
	Version     string    `json:"version"`
	Witnesses   []string  `json:"witnesses"`
}
