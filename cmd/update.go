package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"lukechampine.com/blake3"
)

func CreateCheckSumFile(writePath string, filename string, reader io.ReadCloser) {
	fileExtBlake3 := addBlakeExtension(filename)
	file, _ := os.Create(filepath.Join(writePath, fileExtBlake3))
	var fileBytes []byte
	fileBytes, _ = ioutil.ReadAll(reader)
	byteHash := blake3.Sum256(fileBytes)

	hashReader := bytes.NewReader(byteHash[:])

	if _, err := io.Copy(file, hashReader); err != nil { // that side effect I mentioned in the above for loop makes this possible to do out of loop.
		log.Fatal(err)
	}
}

func isEqualFiles(writePath string, filename string, reader io.ReadCloser) bool {
	fileExtBlake3 := addBlakeExtension(filename)
	previousHash, _ := ioutil.ReadFile(filepath.Join(writePath, fileExtBlake3))
	var fileBytes []byte
	fileBytes, _ = ioutil.ReadAll(reader)
	newHash := blake3.Sum256(fileBytes)
	return bytes.Equal(previousHash, newHash[:])
}

func addBlakeExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos] + ".blake3"
	}
	return fileName + ".blake3"
}
