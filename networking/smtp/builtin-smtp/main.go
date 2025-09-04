package main

import (
	"errors"
	"flag"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	username = flag.String("username", "", "username")
	pass     = flag.String("pass", ".password", "path to password")
	host     = flag.String("host", "smtp.gmail.com", "host")
	addr     = flag.String("addr", "smtp.gmail.com:587", "server addr")
	from     = flag.String("from", "", "from")
	to       = flag.String("to", "", "to")
)

func main() {
	flag.Parse()
	if err := run(*pass, *username, *host, *addr, *from, *to); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("file does not exist: %q\n", *pass)
			os.Exit(1)
		}
	}
}

func run(name string, username string, host string, addr string, from string, to string) error {
	byt, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("error reading file: %q: %w", *pass, err)
	}
	password := string(byt)
	msg := []byte("Subject: Introduction to golang\r\n" + "This is the message of the mail.\r\n")

	auth := smtp.PlainAuth("" /* identity */, username /* username */, password /* password */, host /* host */)
	err = smtp.SendMail(addr /* addr */, auth /* auth */, from /* from */, strings.Split(to, " ") /* to */, msg /* msg */)
	if err != nil {
		return fmt.Errorf("error sending mail: %q: %w", addr, err)
	}
	return nil
}
