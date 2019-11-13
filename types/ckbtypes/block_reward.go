package ckbtypes

type BlockReward struct {
	Total          string `json:"total"`
	Primary        string `json:"primary"`
	Secondary      string `json:"secondary"`
	TxFee          string `json:"tx_fee"`
	ProposalReward string `json:"proposal_reward"`
}
