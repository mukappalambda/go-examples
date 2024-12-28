package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	addr := "localhost:6379"
	readTimeout := 500 * time.Millisecond
	writeTimeout := 500 * time.Millisecond
	opt := &redis.Options{
		Addr:       addr,
		ClientName: "my-redis-client",
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Printf("client is connected: %s\n", cn.ClientGetName(ctx))
			return nil
		},
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	client := redis.NewClient(opt)
	fmt.Printf("client connecting to redis: %s\n", client.Options().Addr)
	defer client.Close()
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("error pinging the redis server: %s\n", err)
	}
	ctx := context.Background()
	cmds, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "name", "alpha", 0)
		pipe.Get(ctx, "name")
		pipe.Set(ctx, "score", 1, 0)
		pipe.Incr(ctx, "score")
		pipe.Get(ctx, "score")
		return nil
	})
	if err != nil {
		client.Close()
		log.Fatalf("error pipelining: %s\n", err)
	}
	fmt.Println(cmds[1].String())
	fmt.Println(cmds[len(cmds)-1].String())
}
