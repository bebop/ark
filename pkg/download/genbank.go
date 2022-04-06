package download

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
)

// Genbank checks the latest release of Genbank, grabs all files ending with .gz extension and saves to disk location specified by writePath.
func Genbank(writePath string) {
	writePathDirectory := filepath.Join(writePath, "genbank")
	links, err := Links("https://ftp.ncbi.nlm.nih.gov/genbank")
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
			go File(link, writePathDirectory)
		}
	}
}
