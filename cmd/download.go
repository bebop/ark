package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download data for standard deploy build",
	Run: func(cmd *cobra.Command, args []string) {
		download()
	},
}

func download() {
	//
}

func getGenbank() error {
	res, err := http.Get("https://ftp.ncbi.nlm.nih.gov/genbank/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Text()
		fmt.Println(title)
		link, _ := s.Attr("href")
		fmt.Println(link)
	})
	return nil
}
