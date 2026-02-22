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
		words := strings.SplitSeq(text, " ")
		for word := range words {
			wordCount[word]++
		}
		fmt.Println("--- word count result ---")
		fmt.Printf("%+v\n", wordCount)
		fmt.Println("--- word count result ---")
	}
}
