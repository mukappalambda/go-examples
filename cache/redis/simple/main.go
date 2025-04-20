package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	addr         = flag.String("addr", "redis:6379", "redis address")
	readTimeout  = flag.Duration("read-timeout", 500*time.Millisecond, "client read timeout")
	writeTimeout = flag.Duration("write-timeout", 500*time.Millisecond, "client write timeout")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	opt := &redis.Options{
		Addr:       *addr,
		ClientName: "my-redis-client",
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Printf("client is connected: %s\n", cn.ClientGetName(ctx))
			return nil
		},
		ReadTimeout:  *readTimeout,
		WriteTimeout: *writeTimeout,
	}
	client := redis.NewClient(opt)
	fmt.Printf("client connecting to redis: %s\n", client.Options().Addr)
	defer client.Close()
	if err := client.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("error pinging the redis server: %w", err)
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
		return fmt.Errorf("error pipelining: %w", err)
	}
	fmt.Println(cmds[1].String())
	fmt.Println(cmds[len(cmds)-1].String())
	return nil
}
