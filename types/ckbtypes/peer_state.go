package ckbtypes

type PeerState struct {
	BlocksInFlight string `json:"blocks_in_flight"`
	LastUpdated    string `json:"last_updated"`
	Peer           string `json:"peer"`
}
