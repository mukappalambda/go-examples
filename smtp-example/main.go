package main

import (
	"flag"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var (
	username = flag.String("username", "", "username")
	pass     = flag.String("pass", "./.password", "path to password")
	host     = flag.String("host", "smtp.gmail.com", "host")
	addr     = flag.String("addr", "smtp.gmail.com:587", "server addr")
	from     = flag.String("from", "", "from")
	to       = flag.String("to", "", "to")
)

func main() {
	flag.Parse()
	byt, err := os.ReadFile(*pass)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("Subject: Introduction to golang\r\n" + "This is the message of the mail.\r\n")

	err = sendMail(*username, string(byt), *host, *addr, *from, strings.Split(*to, " "), msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMail(username string, password string, host string, addr string, from string, to []string, msg []byte) error {
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}
