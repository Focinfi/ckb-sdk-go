package ckbtypes

type LiveCell struct {
	Lock   Script   `json:"lock"`
	Output Output   `json:"output"`
	Data   CellData `json:"data"`
}

type Cell struct {
	BlockHash string   `json:"block_hash"`
	Capacity  string   `json:"capacity"`
	Lock      Script   `json:"lock"`
	OutPoint  OutPoint `json:"out_point"`
}
