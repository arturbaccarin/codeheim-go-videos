package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
	Items       []Item    `xml:"item"`
}

func generateRSSFeed() ([]byte, error) {
	items := []Item{
		{
			Title:       "Article 1",
			Link:        "https://example.com/article1",
			Description: "Description of Article 1",
			PubDate:     time.Now(),
		},
		{
			Title:       "Article 2",
			Link:        "https://example.com/article2",
			Description: "Description of Article 2",
			PubDate:     time.Now().Add(-1 * time.Hour),
		},
	}

	// RSS feed data
	feed := Channel{
		Title:       "Sample RSS Feed",
		Link:        "https://example.com/feed",
		Description: "this is a sample RSS feed generated using Golang.",
		PubDate:     time.Now(),
		Items:       items,
	}

	// Marshal the data to XML
	xmlData, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		return nil, err
	}

	// Add XML header to the feed
	rssFeed := []byte(xml.Header + string(xmlData))

	return rssFeed, nil
}

func main() {
	// Generate the RSS feed
	rssFeed, err := generateRSSFeed()
	if err != nil {
		fmt.Println("error generating rss feed: ", err)
		return
	}

	// Write the feed to a file
	file, err := os.Create("feed.xml")
	if err != nil {
		fmt.Println("error creating file: ", err)
	}
	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		fmt.Println("error writing to file: ", err)
		return
	}

	fmt.Println("rss feed generated successfully")

}
