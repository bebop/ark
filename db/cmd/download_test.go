package cmd

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

/******************************************************************************
So Download as a command is a total pain to test locally and needs a fancy big
server deployment to be fully tested. That's in the works but for now I've
just broken down most of what's needed into helper functions that can be easily
unit tested.

TTFN,
Tim
******************************************************************************/

func TestGetFile(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.RemoveAll(tmpDataDir)

	getFile("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", tmpDataDir)
	getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", tmpDataDir)
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
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.RemoveAll(tmpDataDir)

	response, err := http.Get("https://github.com/TimothyStiles/poly/archive/refs/tags/v0.0.0.tar.gz")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer response.Body.Close()

	err = getTarballFile(response.Body, "README", tmpDataDir)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
