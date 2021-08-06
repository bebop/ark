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

/******************************************************************************
Allbase needs an easy and reproducible way to grab all of the data it uses.
Some people would try to download everything manually with wget or curl but
that's bonkers and I'm not maintaining that.

Instead I've made a multi-threaded webscraper.

It will download and store the latest versions of the following:

	- Rhea RDF file
	- Rhea to Uniprot Sprot mapping file
	- Rhea to Uniprot Trembl mapping file
	- CHEMBL sqlite file
	- Uniprot Sprot XML file
	- Uniprot Trembl XML file
	- All of Genbank's files

Given that these urls are likely not to change I doubt there will be many if
any stability issues from upstream data sources.

TODO: Create flags so that each data source can be downloaded individually.
TODO: Create way to track, report, and resume progress of downloads.

TTFN,
Tim
******************************************************************************/

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download data for standard deploy build. Run at your own risk.",
	Long:  "Download literally downloads all the base data needed to build a standard allbase deployment the amount of data is dummy high to casually test on your personal machine. Run at your own risk.",
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

	// get CHEMBL Sqlite file - ~300MB compressed.
	go getChembl(writePath)

	// get curated sprot uniprot - ~1GB compressed.
	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_sprot.xml.gz", writePath)

	// get chaotic trembl uniprot - ~160GB compressed.
	go getFile("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_trembl.xml.gz", writePath)

	// gets all of annotated genbank - Not sure how big it is as of writing this but it's a lot.
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

// getGenbank checks the latest release of Genbank, grabs all files ending with .gz extension and saves to disk location specified by writePath.
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

		if extension == ".gz" { // if it's a gzipped file it's a genbank file so download it
			fmt.Println("retrieving: " + link)
			go getFile(link, writePathDirectory)
		}
	}
}

// getFile downloads the file at the specified url and saves it to the specified writePath.
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

	// if the filepath does not exist, create it
	err = os.MkdirAll(writePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// parse url for filename
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}

	filename := filepath.Base(parsedURL.Path)
	pathname := filepath.Join(writePath, filename)

	// create a new file to write the data to it
	file, err := os.Create(pathname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// copy the file to disk
	if _, err := io.Copy(file, response.Body); err != nil {
		log.Fatal(err)
	}
}

// getPageLinks returns a slice of all the links on the page at the specified url.
func getPageLinks(url string) ([]string, error) {
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
		if strings.Contains(header.Name, fileNamePattern) { // assuming that our tarball will only contain one file that will match our pattern.
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

		filename += ".gz"

		// create empty file to write to
		file, err := os.Create(filepath.Join(writePath, filename))

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		// compresses the file since it's likely large
		archiver := gzip.NewWriter(file)
		archiver.Name = filename
		defer archiver.Close()

		// copy the compressed file to disk
		if _, err := io.Copy(archiver, directory); err != nil { // that side effect I mentioned in the above for loop makes this possible to do out of loop.
			log.Fatal(err)
		}
	}
	return err
}
