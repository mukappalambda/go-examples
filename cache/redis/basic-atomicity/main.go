package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	redisOptions := &redis.Options{
		Addr: ":6379",
	}
	client := redis.NewClient(redisOptions)
	defer client.Close()
	ctx := context.Background()
	key := "my.key"
	if err := client.Set(ctx, key, 0, 0).Err(); err != nil {
		return fmt.Errorf("failed to set key: %s", err)
	}
	N := 100
	done := make(chan struct{})
	go func() {
		incrN(ctx, client, key, N)
		close(done)
	}()
	<-done
	get := client.Get(ctx, key)
	if err := get.Err(); err != nil {
		return fmt.Errorf("failed to get key: %s", err)
	}
	fmt.Println(key, get.Val())
	return nil
}

func incrN(ctx context.Context, client *redis.Client, key string, N int) {
	_ = client.Watch(ctx, func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			for range N {
				pipe.Incr(ctx, key)
			}
			return nil
		})
		return fmt.Errorf("failed to execute pipeline: %s", err)
	}, key)
}
