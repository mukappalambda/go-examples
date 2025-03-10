package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mukappalambda/go-examples/web/http/http-client/http/core"
)

type remotePosts struct {
	c *Client
}

var _ core.Store = (*remotePosts)(nil)

func (r *remotePosts) Create(ctx context.Context, post *core.Post) (*core.Post, error) {
	panic("unimplemented")
}

func (r *remotePosts) Get(ctx context.Context, id int) (*core.Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/posts/%d", r.c.url, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	post, err := core.PostFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *remotePosts) List(ctx context.Context) ([]*core.Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.c.url+"/posts", nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	posts, err := core.PostsFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (c *Client) PostService() core.Store {
	return &remotePosts{
		c: c,
	}
}
