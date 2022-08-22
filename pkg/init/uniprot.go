package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/uniprot"
	"github.com/uptrace/bun"
)

// Uniprot parses and inserts the uniprot data into the database.
func Uniprot(ctx context.Context, db *bun.DB, config config.Config) error {

	// get uniprot files
	_ = insertUniprotXML(ctx, db, config.UniprotSprotXML)
	if config.UniprotSprotXML != config.UniprotTremblXML { // this only happens in prod
		_ = insertUniprotXML(ctx, db, config.UniprotTremblXML)
	}

	return nil
}

// InsertUniprot inserts the uniprot data into the database.
func insertUniprotXML(ctx context.Context, db *bun.DB, path string) error {
	entries, _, _ := uniprot.Read(path)
	for entry := range entries {
		_, err := db.NewInsert().
			Model(&entry).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}

// CreateUniprotTable creates the uniprot table in the database.
func CreateUniprotTable(ctx context.Context, db *bun.DB) error {

	// create uniprot table
	_, err := db.NewCreateTable().
		Model((*uniprot.Entry)(nil)).
		Exec(ctx)

	return err
}
