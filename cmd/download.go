package cmd

import (
	"archive/tar"
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

	writePath := "../data/build"

	err := getFile("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", writePath)
	if err != nil {
		log.Fatal(err)
	}

	err = getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", writePath)
	if err != nil {
		log.Fatal(err)
	}

	err = getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_trembl.tsv.gz", writePath)
	if err != nil {
		log.Fatal(err)
	}

	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_sprot.xml.gz", writePath)

	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_trembl.xml.gz", writePath)

	go getGenbank()

	go getChembl()

}

func getChembl() error {

	response, err := http.Get("https://ftp.ebi.ac.uk/pub/databases/chembl/ChEMBLdb/latest/chembl_29_sqlite.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// unzip the tarball
	zipReader, err := gzip.NewReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// unpack the tarball
	tarReader := tar.NewReader(zipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Saving " + header.Name + " from CHEMBL")
		if _, err := io.Copy(os.Stdout, tarReader); err != nil {
			log.Fatal(err)
		}
		fmt.Println()

		// create a new file to write the uncompressed data to
		file, err := os.Create(header.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// copy the uncompressed file to disk
		if _, err := io.Copy(file, tarReader); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func getGenbank() error {
	response, err := http.Get("https://ftp.ncbi.nlm.nih.gov/genbank/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Text()
		fmt.Println(title)
		link, _ := s.Attr("href")

		// parse url for file extention
		parsedURL, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}
		filename := filepath.Base(parsedURL.Path)
		extension := filepath.Ext(filename)

		if extension == ".gz" {
			fmt.Println("retrieving: " + link)
			go getFile(link, "../data/build/genbank")
		}
	})
	return nil
}

func getFile(fileURL string, writePath string) error {

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

	err = os.MkdirAll(writePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var pathname string
	if extension == ".gz" {
		pathname = filepath.Join(writePath, filename[0:len(filename)-len(extension)])
	} else {
		pathname = filepath.Join(writePath, filename)
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
