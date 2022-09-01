package init

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/TimothyStiles/allbase/models"
	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/genbank"
	"github.com/uptrace/bun"
)

// Genbank parses and inserts the intitial Genbank data into the database.
func Genbank(ctx context.Context, db *bun.DB, config config.Config) error {

	// get list of Genbank Files
	globPattern := filepath.Join(config.Genbank, "/*.seq.gz")
	genbankFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return err
	}
	if len(genbankFiles) == 0 {
		return errors.New("no Genbank files found")
	}
	// parse each Genbank file
	for _, genbankFile := range genbankFiles {
		// gunzip the file
		zippedFile, err := os.Open(genbankFile)
		if err != nil {
			return err
		}
		defer zippedFile.Close()
		var reader io.Reader
		reader, err = gzip.NewReader(zippedFile)
		if err != nil {
			return err
		}

		// parse the Genbank file into Poly genbank structs
		genbanks, err := genbank.ParseMulti(reader)
		if err != nil {
			return err
		}
		// convert genbanks to genbanksDB
		genbanksDB := make([]models.GenbankWithTags, len(genbanks))
		for i, record := range genbanks {
			genbanksDB[i] = models.GenbankWithTags{
				Meta:     record.Meta,
				Features: record.Features,
				Sequence: record.Sequence,
			}
		}

		// insert the Genbank data into the database (should be changed to upsert)
		_, err = db.NewInsert().
			Model(&genbanksDB).
			Exec(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}

// CreateGenbankTable creates the Genbank table.
func CreateGenbankTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*models.GenbankWithTags)(nil)).
		Exec(ctx)

	return err
}
