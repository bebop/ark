package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/rhea"
	"github.com/TimothyStiles/surrealdb.go"
)

// Rhea parses and inserts the rhea data into the database.
func Rhea(ctx context.Context, db *surrealdb.DB, config config.Config) error {
	// parse Rhea file
	rheaBytes, err := rhea.ReadGzippedXml(config.RheaRDF)
	parsedRhea, err := rhea.Parse(rheaBytes)
	if err != nil {
		return err
	}

	// insert compounds
	for _, compound := range parsedRhea.Compounds {

		// compoundID := "compound:" + strconv.Itoa(compound.ID) // TODO: remove if not needed
		// fmt.Println("parsed rhea", parsedRhea)
		_, err := db.Create("compound", compound)
		if err != nil {
			_, err := db.Update("compound", compound)
			if err != nil {
				return err
			}
		}
	}

	// TODO: insert reaction participants and reaction sides

	// TODO: insert reactions

	return nil
}
