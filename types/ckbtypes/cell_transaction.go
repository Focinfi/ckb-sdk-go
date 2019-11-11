package ckbtypes

type CellTransaction struct {
	ConsumedBy *TransactionPoint `json:"consumed_by,omitempty"`
	CreatedBy  TransactionPoint  `json:"created_by"`
}
