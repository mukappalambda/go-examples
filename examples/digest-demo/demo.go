package main

import (
	"crypto/sha256"
	"fmt"
)

func ReturnSha256sum(data []byte) []byte {
	digest := sha256.Sum256(data)
	return []byte(fmt.Sprintf("%x", digest))
}
