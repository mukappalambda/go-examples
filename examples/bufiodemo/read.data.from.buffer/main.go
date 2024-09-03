package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "你好嗎?🥰"
	buf := bufio.NewReader(strings.NewReader(s))
	b, _ := buf.ReadBytes('\n')
	fmt.Printf("%q\n", b)

	byt := []byte("我很好🤗")
	buf2 := bufio.NewReader(bytes.NewBuffer(byt))
	b2, _ := buf2.ReadBytes('\n')
	fmt.Printf("%q\n", b2)
}
