package download

import (
	"log"
	"net/http"
	"strings"
)

// Chembl checks the latest release for Chembl, downloads and unpacks their sqlite release tarball and saves it to disk write path.
func Chembl(writePath string) {
	links, err := Links("https://ftp.ebi.ac.uk/pub/databases/chembl/ChEMBLdb/latest/")

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
	err = Tarball(response.Body, ".db", writePath)
	if err != nil {
		log.Fatal(err)
	}
}
