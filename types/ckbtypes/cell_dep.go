package ckbtypes

type CellDep struct {
	DepType  DepType  `json:"dep_type"`
	OutPoint OutPoint `json:"out_point"`
}
