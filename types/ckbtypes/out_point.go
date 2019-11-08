package ckbtypes

type OutPoint struct {
	Index  string `json:"index"`
	TxHash string `json:"tx_hash"`
}

func (outPoint OutPoint) Clone() *OutPoint {
	return &OutPoint{
		Index:  outPoint.Index,
		TxHash: outPoint.TxHash,
	}
}
