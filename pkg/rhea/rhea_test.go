package rhea

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bebop/ark/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	testConfig := config.TestDefault()
	var err error
	_, err = Read(testConfig.RheaRDF)
	if err != nil {
		log.Fatalf("Failed to read rhea: %v", err)
	}

	// Start running tests
	code := m.Run()
	os.Exit(code)
}

func ExampleRhea_ExportJSON() {
	// Convert rhea to JSON
	testConfig := config.TestDefault()
	rhea, _ := Read(testConfig.RheaRDF)
	rheaJSON, _ := rhea.ExportJSON()

	fmt.Println(string(rheaJSON)[:100])
	// Output: {"reactionParticipants":[{"compound":"http://rdf.rhea-db.org/Participant_10000_compound_1283","react
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
