package command

import (
	"context"
	"net/http"
	"fmt"
	"io"
	"encoding/xml"
	"html"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting feed: %v", err)
	}
	req.Header.Set("User-Agent", "gator")

	// use the mysterious Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting request: %v", err)
	}
	// check response code for ikk
	if resp.StatusCode > 299 {
		fmt.Println(req)
		return nil, fmt.Errorf("network error: %v", resp.Status)
	}
	// read the response
	data, err := io.ReadAll(resp.Body)

	var feed RSSFeed
	// unmarshall thy data
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("error reading xml into go struct: %v", err)
	}

	// func UnescapeString(s string) string
	fmt.Printf("Description: %v", feed.Channel.Description)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	fmt.Printf("Description: %v", feed.Channel.Description)

	return &feed, nil
}

/*
Use the html.UnescapeString function to decode escaped HTML entities (like &ldquo;). You'll need to run the Title and Description fields (of both the entire channel as well as the items) through this function.



*/