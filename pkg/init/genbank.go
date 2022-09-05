package init

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/genbank"
	"github.com/TimothyStiles/poly/seqhash"
	"github.com/TimothyStiles/surrealdb.go"
)

// Genbank parses and inserts the intitial Genbank data into the database.
func Genbank(ctx context.Context, db *surrealdb.DB, config config.Config) error {

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

		_, err = db.Use("genbank", config.DBName)
		//
		// insert each Genbank into the database
		for _, genbankRecord := range genbanks {

			// hash sequence
			sequenceHash, err := seqhash.Hash(genbankRecord.Sequence, "DNA", genbankRecord.Meta.Locus.Circular, true)
			genbankRecord.Meta.SequenceHash = sequenceHash

			entryID := "genbank:" + sequenceHash
			// insert the Genbank into the database

			_, err = db.Create(entryID, genbankRecord)
			if err != nil {
				_, err = db.Change(entryID, genbankRecord)

			} else {
				return err
			}
		}
	}

	return nil
}
