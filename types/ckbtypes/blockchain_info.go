package ckbtypes

type BlockChainInfo struct {
	Alerts []struct {
		ID          string `json:"id"`
		Message     string `json:"message"`
		NoticeUntil string `json:"notice_until"`
		Priority    string `json:"priority"`
	} `json:"alerts"`
	Chain                  string `json:"chain"`
	Difficulty             string `json:"difficulty"`
	Epoch                  string `json:"epoch"`
	IsInitialBlockDownload bool   `json:"is_initial_block_download"`
	MedianTime             string `json:"median_time"`
}
