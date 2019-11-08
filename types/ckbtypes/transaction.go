package ckbtypes

type Transaction struct {
	CellDeps    []CellDep     `json:"cell_deps"`
	Hash        string        `json:"hash"`
	HeaderDeps  []string      `json:"header_deps"`
	Inputs      []Input       `json:"inputs"`
	Outputs     []Output      `json:"outputs"`
	OutputsData []string      `json:"outputs_data"`
	Version     string        `json:"version"`
	Witnesses   []interface{} `json:"witnesses"`
}

func (transaction Transaction) ToRaw() RawTransaction {
	return RawTransaction{
		CellDeps:    transaction.CellDeps,
		HeaderDeps:  transaction.HeaderDeps,
		Inputs:      transaction.Inputs,
		Outputs:     transaction.Outputs,
		OutputsData: transaction.OutputsData,
		Version:     transaction.Version,
		Witnesses:   transaction.Witnesses,
	}
}

type RawTransaction struct {
	CellDeps    []CellDep     `json:"cell_deps"`
	HeaderDeps  []string      `json:"header_deps"`
	Inputs      []Input       `json:"inputs"`
	Outputs     []Output      `json:"outputs"`
	OutputsData []string      `json:"outputs_data"`
	Version     string        `json:"version"`
	Witnesses   []interface{} `json:"witnesses"`
}
