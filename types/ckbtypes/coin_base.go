package ckbtypes

type CoinBase struct {
	Primary        string `json:"primary"`
	ProposalReward string `json:"proposal_reward"`
	Secondary      string `json:"secondary"`
	Total          string `json:"total"`
	TxFee          string `json:"tx_fee"`
}
