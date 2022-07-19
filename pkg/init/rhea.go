package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/rhea"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Rhea parses and inserts Rhea data into the database.
func Rhea(ctx context.Context, db *sqlx.DB, config config.Config) error {
	// parse Rhea file
	rheaBytes, err := rhea.ReadGzippedXML(config.RheaRDF)
	columns := boil.Infer()
	if err != nil {
		return err
	}

	parsedRhea, err := rhea.Parse(rheaBytes)
	if err != nil {
		return err
	}

	// insert Rhea reactions into database
	for _, reaction := range parsedRhea.Reactions {
		err = reaction.Insert(ctx, db, columns)
		if err != nil {
			// return err
		}
	}

	// insert Rhea compounds into database
	for _, compound := range parsedRhea.Compounds {
		err = compound.Insert(ctx, db, columns)
		if err != nil {
			return err
		}
	}

	// insert Rhea reactionParticipants into database
	for _, reactionParticipant := range parsedRhea.ReactionParticipants {
		err = reactionParticipant.Insert(ctx, db, columns)
		if err != nil {
			return err
		}
	}

	return nil
}
