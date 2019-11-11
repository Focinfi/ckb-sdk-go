package ckbtypes

type CellInfo struct {
	Lock   Script   `json:"lock"`
	Output Output   `json:"output"`
	Data   CellData `json:"data"`
}

type CellOutputWithOutPoint struct {
	BlockHash string   `json:"block_hash"`
	Capacity  string   `json:"capacity"`
	Lock      Script   `json:"lock"`
	OutPoint  OutPoint `json:"out_point"`
}

type LiveCell struct {
	CellOutput Output           `json:"cell_output"`
	CreatedBy  TransactionPoint `json:"created_by"`
}
