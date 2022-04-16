package download

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// File downloads the file at the specified url and saves it to the specified writePath.
func File(fileURL string, writePath string) {
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
