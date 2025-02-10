package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Member struct {
	ID       int
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
	Salary   float64
	Email    string
	Tags     []string
}

func main() {
	f, err := os.Open("./example.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %s\n", err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	lineno := 0
	allMembers := make([]Member, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read line: %s\n", err)
		}

		// ignore line 1
		if lineno == 0 {
			lineno++
			continue
		}

		values := strings.Split(line, ",")
		id, _ := strconv.Atoi(values[0])
		name := values[1]
		age, _ := strconv.Atoi(values[2])
		birthday, _ := time.Parse(time.DateOnly, values[3])
		active, _ := strconv.ParseBool(values[4])
		salary, _ := strconv.ParseFloat(values[5], 64)
		email := values[6]
		tags := strings.Split(values[7], ",")
		member := Member{
			ID:       id,
			Name:     name,
			Age:      age,
			Birthday: birthday,
			Active:   active,
			Salary:   salary,
			Email:    email,
			Tags:     tags,
		}
		allMembers = append(allMembers, member)
	}
	for _, m := range allMembers {
		fmt.Printf("%+v\n", m)
	}
}
