package init

import (
	"context"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/jmoiron/sqlx"
)

// RheaToUniprot adds relations between Rhea and Uniprot.
func RheaToUniprot(ctx context.Context, db *sqlx.DB, config config.Config) error {
	// get rheaToUniprot relations
	rheaToSprot 
}
