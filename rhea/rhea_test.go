package rhea

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
)

var rhea Rhea

func TestMain(m *testing.M) {
	var err error
	rhea, err = Read("data/rhea.rdf.gz")
	if err != nil {
		log.Fatalf("Failed to read rhea: %v", err)
	}

	// Start running tests
	code := m.Run()
	os.Exit(code)
}

func ExampleRhea_ExportJSON() {
	// Convert rhea to JSON
	rheaJson, _ := rhea.ExportJSON()

	fmt.Println(string(rheaJson)[:100])
	// Output: {"reactionParticipants":[{"reactionside":"http://rdf.rhea-db.org/10000_L","contains":1,"containsn":f
}

func TestReadRheaToUniprot(t *testing.T) {
	lines := make(chan RheaToUniprot, 100)
	go ReadRheaToUniprotTrembl("data/rhea2uniprot_sprot_minimized.tsv.gz", lines)

	var line RheaToUniprot
	for l := range lines {
		line = l
	}

	if line.UniprotID != "P06106" {
		log.Fatalf("Got wrong uniprotId. Expected P06106, got %s", line.UniprotID)
	}
}

func ExampleReadRheaToUniprotSprot() {
	lines := make(chan RheaToUniprot, 100)
	go ReadRheaToUniprotSprot("data/rhea2uniprot_sprot_minimized.tsv", lines)

	var line RheaToUniprot
	for l := range lines {
		line = l
	}

	fmt.Println(line)
	// Output: {10048 UN 10048 P06106}
}
