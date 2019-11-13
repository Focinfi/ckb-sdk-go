package ckbtypes

type BlockTemplate struct {
	BytesLimit       string                `json:"bytes_limit"`
	Cellbase         CellbaseTemplate      `json:"cellbase"`
	CompactTarget    string                `json:"compact_target"`
	CurrentTime      string                `json:"current_time"`
	CyclesLimit      string                `json:"cycles_limit"`
	Dao              string                `json:"dao"`
	Epoch            string                `json:"epoch"`
	Number           string                `json:"number"`
	ParentHash       string                `json:"parent_hash"`
	Proposals        []string              `json:"proposals"`
	Transactions     []TransactionTemplate `json:"transactions"`
	Uncles           []UncleTemplate       `json:"uncles"`
	UnclesCountLimit string                `json:"uncles_count_limit"`
	Version          string                `json:"version"`
	WorkID           string                `json:"work_id"`
}
