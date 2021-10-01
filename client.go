package rpc

const Version = "2.1"

type Client struct {
	url           string
	configuration *Configuration
}

func (c *Client) F(name string) *Function {
	return newFunction(name, c)
}

func NewClient(url string, configuration *Configuration) *Client {
	return &Client{
		url:           url,
		configuration: configuration,
	}
}
