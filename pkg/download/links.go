package download

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Links returns a slice of all the links on the page at the specified url.
func Links(url string) ([]string, error) {
	// get the page
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// if server ain't good, bail
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// parse the page into a document goquery can use
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// initialize links slice to hold all the links pulled from following mapping
	var links []string
	doc.Find("a").Each(func(i int, selection *goquery.Selection) { // use a goquery selector to get all links on the page
		// For each item found, get the link
		link, _ := selection.Attr("href")
		if link != "" { // if the link is not empty append it to the slice
			links = append(links, link)
		}
	})
	return links, err
}
