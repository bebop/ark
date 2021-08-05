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

// download literally downloads all the base data needed to build a standard allbase deployment
// the amount of data is dummy high to casually test on your personal machine. Run at your own risk.
func download() {
	writePath := "../data/build"

	// Typically I'd write these functions to return errors but since I'm using go routines
	// the blocking nature of using channels to report errors would either make the
	// concurrency of go routines moot or make it so the returned errors were not returned until
	// all of the go routines were done which in this case kind of makes reporting errors a bit useless.

	// The solution here is that all of the functions called by the go routines will just log fatal errors.

	// I suppose it may be of some use to report when go routines are finished for the user's sake but that isn't a priority for
	// this pull request.

	// get Rhea - relatively small.
	go getFile("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", writePath)

	// get Rhea to curated uniprot mappings - relatively small.
	go getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", writePath)

	// get Rhea to chaotic uniprot mappings - larger than sprot but still relatively small.
	go getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_trembl.tsv.gz", writePath)

	// CHEMBL Sqlite file - ~20GB decompressed. This WILL decompress and save to file.
	go getChembl(writePath)

	// curated uniprot - ~1GB compressed. This WILL decompress and save to file.
	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_sprot.xml.gz", writePath)

	// chaotic uniprot - ~160GB compressed. This WILL decompress and save to file.
	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_trembl.xml.gz", writePath)

	// literally get all of annotated genbank - Not sure how big it is as of writing this but it's a lot.
	go getGenbank(writePath)

}

// getChembl checks the latest release for Chembl, downloads and unpacks their sqlite release tarball and saves it to disk write path.
func getChembl(writePath string) {
	links, err := getPageLinks("https://ftp.ebi.ac.uk/pub/databases/chembl/ChEMBLdb/latest/")

	if err != nil {
		log.Fatal(err)
	}

	var sqliteFileLink string

	// find the sqlite file link
	for _, link := range links {
		// if it's a sqlite tarball save its link
		if strings.Contains(link, "sqlite.tar.gz") {
			sqliteFileLink = link
			break
		}
	}

	// if we didn't find it, bail.
	if sqliteFileLink == "" {
		log.Fatal("could not find sqlite file link")
	}

	// get the tarball from the server that contains the sqlite file
	response, err := http.Get(sqliteFileLink)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// if server ain't good, bail
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// extract our sqlite file from the tarball and write to disk
	err = getTarballFile(response.Body, ".db", writePath)
	if err != nil {
		log.Fatal(err)
	}
}

// getGenbank checks the latest release of Genbank, grabs all files ending with .gz extension. Decompresses them and saves to disk location specified by writePath.
func getGenbank(writePath string) {
	writePathDirectory := filepath.Join(writePath, "genbank")
	links, err := getPageLinks("https://ftp.ncbi.nlm.nih.gov/genbank")
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {

		parsedURL, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}

		filename := filepath.Base(parsedURL.Path)
		extension := filepath.Ext(filename)

		if extension == ".gz" { // if it's a gzipped file it's a genbank file so download and unzip it
			fmt.Println("retrieving: " + link)
			go getFile(link, writePathDirectory)
		}
	}
}

func getFile(fileURL string, writePath string) {
	// get the file from the server
	response, err := http.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// if server ain't good, bail
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// parse url for file extension
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(parsedURL.Path)
	extension := filepath.Ext(filename)

	var reader io.Reader

	// if the file is a gzipped file, decompress and read it else just read it
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

	// if the filepath does not exist, create it
	err = os.MkdirAll(writePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var pathname string
	if extension == ".gz" {
		pathname = filepath.Join(writePath, filename[0:len(filename)-len(extension)]) // trim off the .gz
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
}

func getPageLinks(url string) ([]string, error) {
	// get the page
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var links []string
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the link
		link, _ := selection.Attr("href")
		if link != "" {
			links = append(links, link)
		}
	})
	return links, err
}

// getTarballFile takes a gzipped tarball via Reader and extracts the first file to match fileNamePattern and then writes it to disk at writePath.
func getTarballFile(responseBody io.ReadCloser, fileNamePattern string, writePath string) error {
	// unzip the tarball
	tarball, err := gzip.NewReader(responseBody)
	if err != nil {
		log.Fatal(err)
	}
	defer tarball.Close()
	// create a new tarball reader to iterate through like a directory
	directory := tar.NewReader(tarball)
	var filename string // will save the filename of the file we're writing
	// iterate through the tarball and save the file we're looking for.
	for {
		header, err := directory.Next() // this creates a side effect that we'll exploit outside of this loop to actually save the file
		if err == io.EOF {              // this is the signal that we're done if we haven't already found the file we're looking for
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(header.Name, fileNamePattern) { // assuming that our tarball will only contain one file that will match our patten.
			filename = filepath.Base(header.Name)
			break
		}
	}

	// if the file exists write to disk
	if filename != "" {

		// if the filepath does not exist, create it
		err = os.MkdirAll(writePath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// create empty file to write to
		file, err := os.Create(filepath.Join(writePath, filename))

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		// copy the uncompressed file to disk
		if _, err := io.Copy(file, directory); err != nil { // that side effect I mentioned in the above for loop makes this possible to do out of loop.
			log.Fatal(err)
		}
	}
	return err
}
