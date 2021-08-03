package cmd

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

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

func getFile(fileURL string) error {
	response, err := http.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// parse url for file extention
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(parsedURL.Path)
	extension := filepath.Ext(filename)

	var reader io.Reader

	if extension == ".gz" {
		// open the compressed file
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// open the uncompressed file
		reader = response.Body
	}

	err = os.MkdirAll("../data/build", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var pathname string
	if extension == ".gz" {
		pathname = filepath.Join("../data/build", filename[0:len(filename)-len(extension)])
	} else {
		pathname = filepath.Join("../data/build", filename)
	}
	// create a new file to write the uncompressed data to
	file, err := os.Create(pathname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// copy the uncompressed file to disk
	if _, err := io.Copy(file, reader); err != nil {
		log.Fatal(err)
	}

	return nil
}
