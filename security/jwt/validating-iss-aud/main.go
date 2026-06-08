package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key                = []byte("my-secret-key")
	validSigningMethod = jwt.SigningMethodHS256
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	defaultIssuer := "example.auth.server"
	defaultAudience := "example.api"
	http.HandleFunc("GET /token", HandleToken(defaultIssuer, defaultAudience))
	http.HandleFunc("GET /data", JWTValidator(defaultIssuer, defaultAudience)(HandleData()))
	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 300 * time.Millisecond}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func HandleToken(iss, aud string) http.HandlerFunc {
	method := validSigningMethod
	return func(w http.ResponseWriter, r *http.Request) {
		claims := NewClaims(iss, aud)
		tokenString, err := NewTokenString(method, claims)
		if err != nil {
			http.Error(w, "error generating token", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte(tokenString)); err != nil {
			http.Error(w, "error writing to response", http.StatusInternalServerError)
		}
	}
}

func JWTValidator(iss, aud string) func(http.HandlerFunc) http.HandlerFunc {
	middleware := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				errStr := "unauthorized"
				logger.Error("auth-header", "remote-addr", r.RemoteAddr, "err", errStr)
				http.Error(w, errStr, http.StatusUnauthorized)
				return
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				errStr := "invalid token"
				logger.Error("token", "remote-addr", r.RemoteAddr, "err", errStr)
				http.Error(w, errStr, http.StatusUnauthorized)
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := ParseTokenString(tokenString, iss, aud)
			if err != nil {
				logger.Error("token", "remote-addr", r.RemoteAddr, "err", err.Error())
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			iss, issErr := claims.GetIssuer()
			aud, audErr := claims.GetAudience()
			if issErr != nil || audErr != nil {
				log.Println(issErr, audErr)
				return
			}
			logger.Info("claims", "issuer", iss, "audience", aud)
			h.ServeHTTP(w, r)
		}
	}
	return middleware
}

func HandleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("data")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewClaims(iss, aud string) jwt.Claims {
	claims := jwt.MapClaims{
		"aud": aud,
		"iss": iss,
	}
	return claims
}

func NewTokenString(method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(key)
}

func ParseTokenString(tokenString, iss, aud string) (jwt.MapClaims, error) {
	claims := NewClaims(iss, aud)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithIssuer(iss), jwt.WithAudience(aud))
	if err != nil {
		return nil, err
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, errors.New("invalid token")
	}
	return mapClaims, nil
}
