package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[Error]: %s", err)
		os.Exit(1)
	}
}

func run() error {
	username := flag.String("username", "", "username")
	flag.Parse()
	if len(*username) == 0 {
		flag.PrintDefaults()
		return fmt.Errorf("[run] forgot to pass the `-username` flag?")
	}
	host := "smtp.gmail.com"
	addr := fmt.Sprintf("%s:587", host)
	password_path := os.Getenv("PASSWORD_PATH")
	if len(password_path) == 0 {
		return fmt.Errorf("[run] forgot to pass the `PASSWORD_PATH` environment variable?")
	}

	password, err := readPassword(password_path)
	if err != nil {
		return fmt.Errorf("[run]: failed to read password: %s", err)
	}

	auth := sasl.NewPlainClient("", *username, string(password))

	from := *username
	to := []string{*username}
	msg := strings.NewReader("Subject:My subject\n\nMy primary message goes here")

	if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return fmt.Errorf("[run]: failed to send mail: %s", err)
	}
	return nil
}

func readPassword(name string) (password []byte, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("[readPassword] failed to open \"%s\": %s", name, err)
	}
	defer f.Close()
	password, err = io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("[readPassword] failed to read \"%s\": %s", name, err)
	}
	return password, nil
}
