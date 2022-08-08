package init

import (
	"compress/gzip"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/TimothyStiles/allbase/models"
	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/genbank"
	"github.com/TimothyStiles/poly/seqhash"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Genbank parses and inserts initial Genbank data into the database.
func Genbank(ctx context.Context, db *sqlx.DB, config config.Config) error {

	columns := boil.Infer()
	// parse Genbank files

	// get list of Genbank files
	globPattern := filepath.Join(config.Genbank, "/*.seq.gz")
	genbankFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return err
	}

	// parse Genbank files
	for _, genbankFile := range genbankFiles {
		// gunzip file
		zippedFile, err := os.Open(genbankFile)
		if err != nil {
			return err
		}

		var reader io.Reader
		reader, err = gzip.NewReader(zippedFile)
		if err != nil {
			return err
		}

		// read Genbank file into Poly struct
		genbank, err := genbank.Parse(reader)
		if nil != err {
			// warning: Genbank file could not be parsed
			// logging.Logger(ctx).Warningf("Genbank file could not be parsed: %s", genbankFile)
			continue
		}

		// initalize models.Genbank struct
		var genbankModel models.Genbank
		// populate models.Genbank struct
		genbankModel.Accession = null.StringFrom(genbank.Meta.Accession)
		genbankHash, err := seqhash.Hash(genbank.Sequence, genbank.Meta.Locus.MoleculeType, genbank.Meta.Locus.Circular, true)
		genbankModel.Seqhash = genbankHash
		if err != nil {
			return err
		}
		genbankModel.Insert(ctx, db, columns)
	}

	return nil
}
