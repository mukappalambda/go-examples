package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type CommentRequest struct {
	User   string `json:"user"`
	Public bool   `json:"public"`
	Body   string `json:"body"`
}

func NewCommentRequest(user string, public bool, body string) *CommentRequest {
	return &CommentRequest{User: user, Public: public, Body: body}
}

type CommentResponse struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	Public    bool      `json:"public"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MarshalJSON implements json.Marshaler.
func (c *CommentResponse) MarshalJSON() ([]byte, error) {
	type Tmp struct {
		ID        int    `json:"id"`
		User      string `json:"user"`
		Public    bool   `json:"public"`
		Body      string `json:"body"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	tmp := &Tmp{
		ID:        c.ID,
		User:      c.User,
		Public:    c.Public,
		Body:      c.Body,
		CreatedAt: c.CreatedAt.Format(time.DateTime),
		UpdatedAt: c.UpdatedAt.Format(time.DateTime),
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (c *CommentResponse) UnmarshalJSON(data []byte) error {
	type Tmp struct {
		ID        int    `json:"id"`
		User      string `json:"user"`
		Public    bool   `json:"public"`
		Body      string `json:"body"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	var tmp Tmp
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	createdAt, err := time.Parse(time.DateTime, tmp.CreatedAt)
	if err != nil {
		return err
	}
	updatedAt, err := time.Parse(time.DateTime, tmp.UpdatedAt)
	if err != nil {
		return err
	}
	*c = CommentResponse{
		ID:        tmp.ID,
		User:      tmp.User,
		Public:    tmp.Public,
		Body:      tmp.Body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return nil
}

var (
	_ json.Marshaler   = (*CommentResponse)(nil)
	_ json.Unmarshaler = (*CommentResponse)(nil)
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("POST /comments", createPost())
	ts := httptest.NewServer(nil)
	defer ts.Close()
	commentReq := NewCommentRequest("alpha", false, "my-comment")
	b, err := json.Marshal(commentReq)
	if err != nil {
		return fmt.Errorf("could not marshal comment: %s", err)
	}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/comments", ts.URL), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to create new request: %s", err)
	}
	client := ts.Client()
	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer res.Body.Close()
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err)
	}
	var commentRes CommentResponse
	if err := json.Unmarshal(b, &commentRes); err != nil {
		return fmt.Errorf("error unmarshaling response: %s", err)
	}
	fmt.Printf("%+v\n", commentRes)
	return nil
}

func createPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		buf, err := io.ReadAll(body)
		if err != nil {
			log.Printf("error reading request body: %s\n", err)
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}
		var commentRequest CommentRequest
		err = json.Unmarshal(buf, &commentRequest)
		if err != nil {
			log.Printf("error unmarshaling request body: %s\n", err)
			http.Error(w, "error  request body", http.StatusBadRequest)
			return
		}
		log.Printf("[client@%s] created comment: %+v\n", r.RemoteAddr, commentRequest)
		createAt := time.Now()
		commentResponse := &CommentResponse{
			ID:        1,
			User:      commentRequest.User,
			Public:    commentRequest.Public,
			Body:      commentRequest.Body,
			CreatedAt: createAt,
			UpdatedAt: createAt,
		}
		b, err := json.Marshal(commentResponse)
		if err != nil {
			log.Printf("error marshaling resource: %s\n", err)
			http.Error(w, "error marshaling resource", http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprint(w, string(b)); err != nil {
			log.Printf("error writing response: %s\n", err)
			http.Error(w, "error writing response", http.StatusInternalServerError)
			return
		}
	}
}
