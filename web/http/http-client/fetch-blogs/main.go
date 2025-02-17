package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	timeout := 3 * time.Second
	client := &http.Client{Timeout: timeout}
	url := `https://go.dev/blog/all`
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create new request: %s", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("failed to create document from response body: %s", err)
	}
	titles := fetchBlogTitles(doc)
	summaries := fetchBlogSummaries(doc)

	for i := 0; i < len(titles); i++ {
		fmt.Printf("title: %q; summary: %q\n", titles[i], summaries[i])
	}
	return nil
}

func fetchBlogTitles(doc *goquery.Document) []string {
	titles := make([]string, 0)
	doc.Find(".blogtitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		titles = append(titles, title)
	})
	return titles
}

func fetchBlogSummaries(doc *goquery.Document) []string {
	summaries := make([]string, 0)
	doc.Find(".blogsummary").Each(func(i int, s *goquery.Selection) {
		summary := strings.TrimSpace(s.Text())
		summaries = append(summaries, summary)
	})
	return summaries
}
