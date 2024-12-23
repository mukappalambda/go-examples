package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wordCount := make(map[string]int)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")
		for _, word := range words {
			wordCount[word] = wordCount[word] + 1
		}
		fmt.Println("--- word count result ---")
		fmt.Printf("%+v\n", wordCount)
		fmt.Println("--- word count result ---")
	}
}
