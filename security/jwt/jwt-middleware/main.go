package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey   []byte
	validMethod = jwt.SigningMethodHS256
)

func main() {
	secretKey = []byte("secret key")
	addr := ":8080"
	http.Handle("GET /token", HandleToken())
	http.Handle("GET /data", JWTMiddleware(HandleData()))
	var wg sync.WaitGroup
	wg.Go(func() {
		server := &http.Server{Addr: addr, ReadHeaderTimeout: 300 * time.Millisecond}
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	})
	fmt.Printf("Server listening on %s\n", addr)
	wg.Wait()
}

func HandleToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := GenerateTokenString(secretKey)
		if err != nil {
			http.Error(w, "error generating token string", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	}
}

func JWTMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerString := r.Header.Get("Authorization")
		if headerString == "" {
			http.Error(w, "invalid authoriztion", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(headerString, "Bearer ") {
			http.Error(w, "invalid authoriztion", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(headerString, "Bearer ")
		token, err := ParseAndValidateTokenString(tokenString)
		if err != nil {
			http.Error(w, "invalid parsing token string", http.StatusUnauthorized)
			return
		}
		claims := token.Claims
		ctx := r.Context()
		type key string
		ctx = context.WithValue(ctx, key("claims"), claims)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HandleData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		claims, ok := ctx.Value("claims").(jwt.MapClaims)
		if !ok {
			http.Error(w, "error converting claims", http.StatusInternalServerError)
			return
		}
		msg, ok := claims["msg"].(string)
		if !ok {
			http.Error(w, "error converting msg type", http.StatusInternalServerError)
		}
		if _, err := w.Write([]byte(msg)); err != nil {
			http.Error(w, "error writing to response", http.StatusInternalServerError)
		}
	})
}

func GenerateTokenString(key []byte) (string, error) {
	token := jwt.NewWithClaims(validMethod, jwt.MapClaims{
		"iss": "my.server",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"msg": "extra message",
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing string: %s", err)
	}
	return tokenString, nil
}

func ParseAndValidateTokenString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("error parsing claims")
		}
		return secretKey, nil
	}, jwt.WithIssuer("my.server"))
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
