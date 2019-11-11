package rpc

const DefaultURL = "http://localhost:8114"

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	if url == "" {
		url = DefaultURL
	}
	return &Client{URL: url}
}
