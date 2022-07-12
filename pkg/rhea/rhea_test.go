package rhea

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

var rhea Rhea

func TestMain(m *testing.M) {
	testConfig := config.TestDefault()
	var err error
	rhea, err = Read(testConfig.RheaRDF)
	if err != nil {
		log.Fatalf("Failed to read rhea: %v", err)
	}

	// Start running tests
	code := m.Run()
	os.Exit(code)
}

func ExampleRhea_ExportJSON() {
	// Convert rhea to JSON
	rheaJSON, _ := rhea.ExportJSON()

	fmt.Println(string(rheaJSON)[:100])
	// Output: {"reactionParticipants":[{"reactionside":"http://rdf.rhea-db.org/10000_L","contains":1,"containsn":f
}

func TestReadRheaToUniprot(t *testing.T) {
	testConfig := config.TestDefault()
	lines := make(chan RheaToUniprot, 100)
	go ReadRheaToUniprotTrembl(testConfig.RheaToUniprotSprot, lines)

	var line RheaToUniprot
	for l := range lines {
		line = l
	}

	if line.UniprotID != "P06106" {
		log.Fatalf("Got wrong uniprotId. Expected P06106, got %s", line.UniprotID)
	}
}

func ExampleReadRheaToUniprotSprot() {
	testConfig := config.TestDefault()
	lines := make(chan RheaToUniprot, 100)
	go ReadRheaToUniprotSprot(testConfig.RheaToUniprotSprot, lines)

	var line RheaToUniprot
	for l := range lines {
		line = l
	}

	fmt.Println(line)
	// Output: {10048 UN 10048 P06106}
}
