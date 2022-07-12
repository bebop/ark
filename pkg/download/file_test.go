package download_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/download"
)

func TestFile(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.RemoveAll(tmpDataDir)

	download.File("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", tmpDataDir)
	download.File("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", tmpDataDir)
}
