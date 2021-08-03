package cmd

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

// Rhea:
// https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz

// Rhea TSV sprot:
// https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot%5Fsprot.tsv

// Rhea TSV trembl:
// https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot%5Ftrembl.tsv.gz

// Uniprot sprot:
// https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_sprot.xml.gz

// Uniprot trembl:
// https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_trembl.xml.gz

func getRhea() error {
	// get the compressed rhea file
	res, err := http.Get("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz")
	if err != nil {
		log.Fatal(err)
	}

	// not sure why Golang http stdlib needs to close this but here we are
	defer res.Body.Close()

	// if the status code is whack, bail
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// open the compressed file
	r, err := gzip.NewReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	err = os.MkdirAll("../data/build", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create a new file to write the uncompressed data to
	f, err := os.Create("../data/build/rhea.rdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// copy the uncompressed file to disk
	if _, err := io.Copy(f, r); err != nil {
		log.Fatal(err)
	}
	return nil
}
