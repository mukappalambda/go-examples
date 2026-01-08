package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type MyServer struct {
	conf   *oauth2.Config
	server *http.Server
}

type LoginHandler struct {
	conf *oauth2.Config
}

type CallbackHandler struct {
	conf *oauth2.Config
}

var (
	port        = flag.Uint("port", 8080, "server port")
	readTimeout = flag.Duration("read-timeout", 1*time.Second, "server read timeout")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		fmt.Println("Missing CLIENT_ID or CLIENT_SECRET")
		os.Exit(1)
	}

	redirectURL := fmt.Sprintf("http://localhost:%d/callback", *port)
	conf := newOauth2Config(clientID, clientSecret, redirectURL)
	myServer := newMyServer(addr, conf, *readTimeout)
	fmt.Printf("Server running at %s\nVisit http://localhost:%d/login\n", addr, *port)
	if err := myServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func newMyServer(addr string, conf *oauth2.Config, readTimeout time.Duration) *MyServer {
	myServer := &MyServer{conf: conf, server: &http.Server{Addr: addr, ReadTimeout: readTimeout}}
	setupRoutes(myServer)
	return myServer
}

func setupRoutes(s *MyServer) {
	handler := http.DefaultServeMux
	handler.Handle("GET /login", &LoginHandler{conf: s.conf})
	handler.Handle("GET /callback", &CallbackHandler{conf: s.conf})
	s.server.Handler = handler
}

func (s *MyServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func newOauth2Config(clientID, clientSecret, redirectURL string) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       []string{"repo", "user"},
	}
	return conf
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	verifier := oauth2.GenerateVerifier()
	url := h.conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.VerifierOption(verifier))
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	var err error
	token, e := h.conf.Exchange(ctx, code)
	err = errors.Join(err, e)
	log.Printf("access token: %s\n", token.AccessToken)
	client := h.conf.Client(ctx, token)
	url := "https://api.github.com/repos/mukappalambda/rust-examples/events"
	req, e := http.NewRequest(http.MethodGet, url, nil)
	err = errors.Join(err, e)
	req.Header.Set("Accept", "application/vnd.github+json")
	res, err := client.Do(req)
	err = errors.Join(err, e)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	err = errors.Join(err, e)
	if err != nil {
		res.Body.Close()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
