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
	"strings"

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

	//  get list of latest chembl data
	response, err := http.Get("https://ftp.ebi.ac.uk/pub/databases/chembl/ChEMBLdb/latest/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// load the chembl list so we can parse it to find the latest Sqlite version
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over the list of chembl releases to find the latest Sqlite version
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {

		// get element link
		link, _ := selection.Attr("href")

		// if it's a sqlite file, download it
		if strings.Contains(link, "sqlite.tar.gz") {
			fmt.Println("retrieving: " + link)
			chembl, err := http.Get(link)
			if err != nil {
				log.Fatal(err)
			}
			defer chembl.Body.Close()

			// if chembl server ain't good, bail
			if chembl.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", chembl.StatusCode, chembl.Status)
			}

			// unzip tarball
			gz, err := gzip.NewReader(chembl.Body)
			if err != nil {
				log.Fatal(err)
			}
			defer gz.Close()

			// unpack tarball
			tarReader := tar.NewReader(gz)

			// iterate over the files in the tarball. If it's a sqlite file, save to disk, and break.
			for {

				header, err := tarReader.Next()

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				// If it's a sqlite file, save to disk, and break
				if filepath.Ext(header.Name) == ".db" {

					// create a new file to write the uncompressed data to
					file, err := os.Create(filepath.Join("../data/build", header.Name))
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					// copy the uncompressed file to disk
					if _, err := io.Copy(file, tarReader); err != nil {
						log.Fatal(err)
					}
					break
				}
			}

		}

	})
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
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the link
		link, _ := selection.Attr("href")

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
