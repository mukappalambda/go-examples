package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("my-secret-key")

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"key": "value",
		},
	)
	s, err := token.SignedString(key)
	if err != nil {
		log.Fatalf("Error signing string: %s\n", err)
	}
	fmt.Printf("Signed JWT: %s\n", s)
	secondToken, err := jwt.Parse(s, myKeyFunc)
	if err != nil {
		log.Fatalf("Error parsing the signature: %s\n", err)
	}
	if claims, ok := secondToken.Claims.(jwt.MapClaims); ok && secondToken.Valid {
		fmt.Printf("Claims: %+v\n", claims)
	}
}

func myKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Incorrect signing method")
	}
	return key, nil
}
