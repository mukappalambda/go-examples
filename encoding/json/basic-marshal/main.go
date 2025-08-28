package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type S struct {
	ID      string  `json:"id"`
	Score   float64 `json:"score"`
	IsValid bool    `json:"is_valid"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	vs := []any{
		true,
		123,
		"foobar",
		[]string{"alpha", "beta", "gamma"},
		map[string]float32{"alpha": 1.11, "beta": 2.22, "gamma": 3.33},
		S{
			ID:      "123",
			Score:   12.34,
			IsValid: false,
		},
	}
	var dataSlice [][]byte
	for _, v := range vs {
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal: %s", err)
		}
		dataSlice = append(dataSlice, data)
	}
	for _, data := range dataSlice {
		fmt.Println("data:", string(data))
	}
	return nil
}
