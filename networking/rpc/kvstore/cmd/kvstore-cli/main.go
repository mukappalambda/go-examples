package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/mukappalambda/go-examples/networking/rpc/kvstore/client"
)

var (
	key = flag.String("key", "", "key")
	val = flag.String("value", "", "value")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	if *key == "" {
		flag.PrintDefaults()
		return errors.New("key is missing")
	}
	address := "127.0.0.1:8080"
	c, err := client.NewClient(address)
	if err != nil {
		return err
	}
	defer c.Close()
	if *val != "" {
		if err := c.Set(*key, *val); err != nil {
			return err
		}
		fmt.Println("Set key")
		return nil
	}
	value, err := c.Get(*key)
	if err != nil {
		return err
	}
	fmt.Printf("%q: %q\n", *key, value)
	return nil
}
