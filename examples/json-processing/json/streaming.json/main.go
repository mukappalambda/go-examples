package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Name  string
	Score int
}

func main() {
	fmt.Println("Type something!\nExample:\n> alpha 10\n------")
	reader := bufio.NewReader(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("error reading: %s\n", err)
			continue
		}
		trimmedLine := strings.TrimSuffix(string(line), "\n")
		splitted := strings.Split(trimmedLine, " ")
		if len(splitted) < 2 {
			log.Printf("not enough fields; got %d field only\n", len(splitted))
			continue
		}
		name := splitted[0]
		score, err := strconv.ParseInt(splitted[1], 10, 64)
		if err != nil {
			log.Printf("error parsing %q to integer\n", splitted[1])
			continue
		}
		data := &Data{
			Name:  name,
			Score: int(score),
		}
		if err := enc.Encode(data); err != nil {
			log.Printf("error encoding data: %s\n", err)
			continue
		}
	}
}
