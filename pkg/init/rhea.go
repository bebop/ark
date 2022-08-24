package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/rhea"
	"github.com/uptrace/bun"
)

// Rhea parses and inserts the rhea data into the database.
func Rhea(ctx context.Context, db *bun.DB, config config.Config) error {
	// parse Rhea file
	rheaBytes, err := rhea.ReadGzippedXml(config.RheaRDF)
	if err != nil {
		return err
	}

	parsedRhea, err := rhea.Parse(rheaBytes)
	if err != nil {
		return err
	}

	// insert Rhea into the database
	_, err = db.NewInsert().
		Model(&parsedRhea).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

// CreateRheaTable creates the rhea table in the database.
func CreateRheaTable(ctx context.Context, db *bun.DB) error {

	// create uniprot table
	_, err := db.NewCreateTable().
		Model((*rhea.Rhea)(nil)).
		Exec(ctx)

	return err
}
