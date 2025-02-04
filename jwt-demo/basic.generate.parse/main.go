package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key                      = []byte("my-secret-key")
	defaultExpirationTimeout = 1 * time.Hour
)

var (
	iss         = flag.String("iss", "my.auth.server", "issuer")
	expIss      = flag.String("exp.iss", "my.auth.server", "expected issuer")
	expRequired = flag.Bool("require.exp", false, "requires exp claim. default to false")
)

func main() {
	flag.Parse()

	validSigningMethod := jwt.SigningMethodHS256

	firstToken := jwt.New(validSigningMethod)
	tokenString, err := firstToken.SignedString(key)
	if err != nil {
		log.Fatalf("Error signing token string: %s\n", err)
	}
	fmt.Printf("Signed first token successfully, signed token string: %q\n", tokenString)

	secondToken := jwt.NewWithClaims(validSigningMethod, newMapClaims(*expRequired))
	tokenString, err = secondToken.SignedString(key)
	if err != nil {
		log.Fatalf("Error signing string: %s\n", err)
	}
	fmt.Printf("Signed second token successfully, signed token string: %q\n", tokenString)
	opts := []jwt.ParserOption{
		jwt.WithIssuer(*expIss),
	}
	if *expRequired {
		opts = append(opts, jwt.WithExpirationRequired())
	}
	secondClaims, err := parseTokenString(tokenString, opts...)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range *secondClaims {
		fmt.Printf("key: %q; value: %q\n", k, v)
	}
}

func newMapClaims(expRequired bool) jwt.MapClaims {
	if expRequired {
		expiration := time.Now().Add(defaultExpirationTimeout).Unix()
		claims := jwt.MapClaims{
			"iss":   *iss,
			"exp":   expiration,
			"scope": "read:messages write:messages",
			"roles": "admin user",
			"key":   "value",
		}
		return claims
	}
	claims := jwt.MapClaims{
		"iss":   *iss,
		"scope": "read:messages write:messages",
		"roles": "admin user",
		"key":   "value",
	}
	return claims
}
func parseTokenString(tokenString string, opts ...jwt.ParserOption) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, myKeyFunc, opts...)
	if err != nil {
		return nil, fmt.Errorf("error parsing the token string: %s", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("error converting to custom claims")
	}
	return &claims, nil
}

func myKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("incorrect signing method")
	}
	return key, nil
}
