package init

import (
	"context"

	"github.com/TimothyStiles/allbase/models"
	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/poly/io/uniprot"
	"github.com/TimothyStiles/poly/seqhash"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Uniprot is the function that initializes the uniprot portion of the database.
func Uniprot(ctx context.Context, db *sqlx.DB, config config.Config) error {

	_ = insert(ctx, db, config.UniprotSprotXML)
	if config.UniprotSprotXML != config.UniprotTremblXML { // this only happens in prod
		_ = insert(ctx, db, config.UniprotTremblXML)
	}

	return nil
}

func insert(ctx context.Context, db *sqlx.DB, path string) error {
	entries, _, _ := uniprot.Read(path)

	for xmlUniprot := range entries {
		dbUniprot := models.Uniprot{}
		dbUniprot.Accession = null.StringFrom(xmlUniprot.Accession[0])
		var err error
		dbUniprot.Seqhash, err = seqhash.Hash(xmlUniprot.Sequence.Value, "PROTEIN", false, false)
		if err != nil {
			return err
		}
		dbUniprot.Insert(ctx, db, boil.Columns{})
	}
	return nil
}
