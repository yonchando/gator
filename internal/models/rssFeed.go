package models

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	rssFeed := RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, strings.NewReader(""))

	if err != nil {
		return &rssFeed, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	var res *http.Response

	res, err = client.Do(req)

	if err != nil {
		return &rssFeed, err
	}

	var body []byte

	body, err = io.ReadAll(res.Body)

	if err != nil {
		return &rssFeed, err
	}

	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		return &RSSFeed{}, err
	}

	return &rssFeed, nil
}
