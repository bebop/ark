package download

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Tarball takes a gzipped tarball via Reader and extracts the first file to match fileNamePattern and then writes it to disk at writePath.
func Tarball(responseBody io.ReadCloser, fileNamePattern string, writePath string) error {
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

		// create empty file to write to
		file, err := os.Create(filepath.Join(writePath, filename))

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		// copy the compressed file to disk
		if _, err := io.Copy(file, directory); err != nil { // that side effect I mentioned in the above for loop makes this possible to do out of loop.
			log.Fatal(err)
		}
	}
	return err
}
