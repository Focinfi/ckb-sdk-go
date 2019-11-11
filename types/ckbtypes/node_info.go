package ckbtypes

type NodeInfo struct {
	Addresses []struct {
		Address string `json:"address"`
		Score   string `json:"score"`
	} `json:"addresses"`
	IsOutbound *bool  `json:"is_outbound"`
	NodeID     string `json:"node_id"`
	Version    string `json:"version"`
}
