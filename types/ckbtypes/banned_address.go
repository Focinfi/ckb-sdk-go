package ckbtypes

type BannedAddress struct {
	Address   string `json:"address"`
	BanReason string `json:"ban_reason"`
	BanUntil  string `json:"ban_until"`
	CreatedAt string `json:"created_at"`
}
