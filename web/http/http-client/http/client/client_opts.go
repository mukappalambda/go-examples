package client

import "time"

type clientOpts struct {
	url     string
	timeout time.Duration
}

type Opt func(*clientOpts)

func WithDefaultTimeout(d time.Duration) Opt {
	return func(c *clientOpts) {
		c.timeout = d
	}
}
