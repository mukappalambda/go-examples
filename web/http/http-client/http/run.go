package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/mukappalambda/go-examples/web/http/http-client/http/client"
)

func run() error {
	var (
		get     = flag.Bool("get", false, "get post by id")
		create  = flag.Bool("create", false, "create post")
		list    = flag.Bool("list", false, "list posts")
		id      = flag.Int("id", 0, "post id")
		timeout = flag.Duration("timeout", 5*time.Second, "client timeout")
	)

	flag.Parse()
	if !(*list) && !(*get) && !(*create) {
		flag.PrintDefaults()
		return fmt.Errorf("specify at least one of operations")
	}
	copts := []client.Opt{
		client.WithDefaultTimeout(*timeout),
	}
	client := client.New(copts...)

	if *list {
		posts, err := client.PostService().List(context.Background())
		if err != nil {
			return fmt.Errorf("failed to list posts: %w", err)
		}

		for _, post := range posts {
			post.Display()
		}
		return nil
	}

	if *get {
		if *id < 1 {
			return fmt.Errorf("invalid id")
		}
		post, err := client.PostService().Get(context.Background(), *id)
		if err != nil {
			return fmt.Errorf("failed to get post: %w", err)
		}
		post.Display()
		return nil
	}

	if *create {
		return fmt.Errorf("unimplemented")
		// // post request
		// jsonStr := []byte(`{"userId": 1, "title": "post title", "body": "post body"}`)
		// req, err = http.NewRequestWithContext(context.Background(), http.MethodPost, *url, bytes.NewBuffer(jsonStr))
		// if err != nil {
		// 	return fmt.Errorf("making post request: %q; %w", *url, err)
		// }

		// req.Header.Set("Content-Type", "application/json")

		// resp, err = client.Do(req)

		// if err != nil {
		// 	return fmt.Errorf("retrieving response: %w", err)
		// }

		// defer resp.Body.Close()
		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	return fmt.Errorf("reading response body: %w", err)
		// }
		// fmt.Println(string(body))
	}
	return nil
}
