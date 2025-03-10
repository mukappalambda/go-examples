package client

import "net/http"

type Client struct {
	url        string
	httpClient *http.Client
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

func New(opt ...Opt) *Client {
	var copts clientOpts
	copts.url = url
	for _, o := range opt {
		o(&copts)
	}
	if copts.timeout == 0 {
		copts.timeout = defaultClientTimeout
	}
	httpClient := &http.Client{
		Timeout: copts.timeout,
	}
	return &Client{
		url:        copts.url,
		httpClient: httpClient,
	}
}
