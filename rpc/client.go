package rpc

import "net/http"

type Client struct {
	url        string
	httpClient *http.Client
}

func (c *Client) F(name string) *Function {
	return newFunction(name, c)
}

func NewClient(url string) *Client {
	return &Client{
		url:        url,
		httpClient: &http.Client{},
	}
}
