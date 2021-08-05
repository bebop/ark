package cmd

import (
	"net/http"
	"testing"
)

func TestGetGenbank(t *testing.T) {
	// getGenbank()
}
func TestGetFile(t *testing.T) {
	getFile("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", "../data/build")
	getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", "../data/build")
}

func TestGetPageLinks(t *testing.T) {
	links, err := getPageLinks("http://example.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(links) == 0 {
		t.Errorf("Error: No links found")
	}
}

func TestGetTarballFile(t *testing.T) {
	response, err := http.Get("https://github.com/TimothyStiles/poly/archive/refs/tags/v0.0.0.tar.gz")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = getTarballFile(response.Body, "README", "../data/test")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
