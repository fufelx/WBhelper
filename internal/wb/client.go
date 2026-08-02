package wb

import (
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	token   string
	http    *http.Client
}

func New(token string) *Client {
	return &Client{
		baseURL: "https://common-api.wildberries.ru",
		token:   token,
		http:    &http.Client{Timeout: 10 * time.Second},
	}
}
