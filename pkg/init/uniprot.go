package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/uniprot"
	"github.com/TimothyStiles/poly/seqhash"
	"github.com/TimothyStiles/surrealdb.go"
)

// Uniprot parses and inserts the uniprot data into the database.
func Uniprot(ctx context.Context, db *surrealdb.DB, config config.Config) error {

	// get uniprot files
	_ = insertUniprotXML(ctx, db, config.UniprotSprotXML)
	if config.UniprotSprotXML != config.UniprotTremblXML { // this only happens in prod
		_ = insertUniprotXML(ctx, db, config.UniprotTremblXML)
	}

	return nil
}

// InsertUniprot inserts the uniprot data into the database.
func insertUniprotXML(ctx context.Context, db *surrealdb.DB, path string) error {
	entries, _, _ := uniprot.Read(path)
	for entry := range entries {
		// insert uniprot entry
		sequenceHash, err := seqhash.Hash(entry.Sequence.Value, "PROTEIN", false, false)
		entryID := "uniprot:" + sequenceHash

		_, err = db.Create(entryID, entry)
		if err != nil {
			_, err = db.Change(entryID, entry)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
