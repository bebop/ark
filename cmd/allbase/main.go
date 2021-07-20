package main

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/TimothyStiles/poly"
	"github.com/TimothyStiles/poly/parsers/uniprot"
	"github.com/allyourbasepair/allbase"
	"github.com/allyourbasepair/allbase/rhea"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Begin SQLite
	db, err := sqlx.Open("sqlite3", "allbase.db")
	if err != nil {
		log.Fatalf("Failed to open sqlite in allbase.db: %s", err)
	}

	// Execute our schema in memory
	_, err = db.Exec(allbase.Schema)
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	// Insert Rhea
	log.Printf("Inserting rhea")
	rhea, err := rhea.Read("rhea/data/rhea_mini.rdf.gz")
	if err != nil {
		log.Fatalf("Could not read rhea: %s", err)
	}

	err = allbase.RheaInsert(db, rhea)
	if err != nil {
		log.Fatalf("Could not insert rhea: %s", err)
	}

	// Insert Uniprot
	log.Printf("Inserting uniprot sprot")
	var wg sync.WaitGroup
	uniprotSprot, errors, err := uniprot.Read("uniprot/uniprot_sprot.xml.gz")
	if err != nil {
		log.Fatalf("Failed to read uniprot on error: %s", err)
	}
	wg.Add(1)
	go allbase.UniprotInsert(db, "sprot", uniprotSprot, errors, &wg)
	wg.Wait()

	for err := range errors {
		if err.Error() != "EOF" {
			log.Fatalf("Failed on error during uniprot parsing or insertion: %s", err)
		}
	}

	log.Printf("Inserting uniprot trembl")
	var wg2 sync.WaitGroup
	uniprotTrembl, errors, err := uniprot.Read("uniprot/uniprot_trembl.xml.gz")
	if err != nil {
		log.Fatalf("Failed to read uniprot on error: %s", err)
	}
	wg2.Add(1)
	go allbase.UniprotInsert(db, "trembl", uniprotTrembl, errors, &wg2)
	wg2.Wait()

	for err := range errors {
		if err.Error() != "EOF" {
			log.Fatalf("Failed on error during uniprot trembl parsing or insertion: %s", err)
		}
	}

	// Insert Genbank
	matches, err := filepath.Glob("genbank/*")
	if err != nil {
		log.Fatalf("Failed during opening glob: %s", err)
	}
	for _, match := range matches {
		log.Printf("Inserting genbank file %s", match)
		sequences := poly.ReadGbkFlatGz("data/flatGbk_test.seq.gz")
		err := allbase.GenbankInsert(db, sequences)
		if err != nil {
			log.Fatalf("Failed on error during genbank insertion: %s", err)
		}
	}

	// Insert tsv
	log.Printf("Inserting rhea->uniprot sprot")
	err = allbase.RheaTsvInsert(db, "rhea/rhea2uniprot_sprot.tsv", false)
	if err != nil {
		log.Fatalf("Failed to insert RheaTsvInsert sprot on: %s", err)
	}

	log.Printf("Inserting rhea->uniprot trembl")
	err = allbase.RheaTsvInsert(db, "rhea/rhea2uniprot_trembl.tsv.gz", true)
	if err != nil {
		log.Fatalf("Failed to insert RheaTsvInsert trembl on: %s", err)
	}

	log.Printf("Finished allbase")
}
