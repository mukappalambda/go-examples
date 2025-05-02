package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/redis/go-redis/v9"
)

var cli struct {
	Create CreateCmd `cmd:"" aliases:"c" help:"Create a key value pair"`
	Read   ReadCmd   `cmd:"" aliases:"r" help:"Read a key"`
	Update UpdateCmd `cmd:"" aliases:"u" help:"Update a key value pair"`
	Delete DeleteCmd `cmd:"" aliases:"d" help:"Delete a key"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	opts := []kong.Option{
		kong.Name("crudctl"),
		kong.Description("rgctl is a CLI tool that interacts with the Redis instance"),
		kong.UsageOnError(),
	}
	return kong.Parse(&cli, opts...).Run()
}

var (
	redisOptions = &redis.Options{
		Addr:                  ":6379",
		DB:                    0,
		MaxRetries:            -1,
		DialTimeout:           10 * time.Second,
		ReadTimeout:           30 * time.Second,
		WriteTimeout:          30 * time.Second,
		ContextTimeoutEnabled: true,
	}
	client = redis.NewClient(redisOptions)
)

func (c *CreateCmd) Run() error {
	if err := client.Set(context.Background(), c.Key, c.Value, 0).Err(); err != nil {
		return err
	}
	fmt.Printf("Created [Key]: %s [Value]: %s\n", c.Key, c.Value)
	return nil
}

func (r *ReadCmd) Run() error {
	get := client.Get(context.Background(), r.Key)
	if err := get.Err(); err != nil {
		return err
	}
	fmt.Printf("[Key]: %s [Value]: %s\n", r.Key, get.Val())
	return nil
}

func (u *UpdateCmd) Run() error {
	if err := client.Set(context.Background(), u.Key, u.Value, 0).Err(); err != nil {
		return err
	}
	fmt.Printf("Updated [Key]: %s [Value]: %s\n", u.Key, u.Value)
	return nil
}

func (r *DeleteCmd) Run() error {
	del := client.Del(context.Background(), r.Key)
	if err := del.Err(); err != nil {
		return err
	}
	fmt.Printf("Deleted [Key]: %s\n", r.Key)
	return nil
}

type CreateCmd struct {
	Key   string `short:"k" required:"" help:"key"`
	Value string `short:"v" required:"" help:"value"`
}

type ReadCmd struct {
	Key string `short:"k" required:"" help:"key"`
}

type UpdateCmd struct {
	Key   string `short:"k" required:"" help:"key"`
	Value string `short:"v" required:"" help:"value"`
}

type DeleteCmd struct {
	Key string `short:"k" required:"" help:"key"`
}
