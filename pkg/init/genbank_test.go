package init

import (
	"context"
	"database/sql"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func TestGenbank(t *testing.T) {

	ctx := context.Background()
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	err = CreateGenbankTable(ctx, db)
	if err != nil {
		panic(err)
	}

	type args struct {
		ctx    context.Context
		db     *bun.DB
		config config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestGenbank",
			args: args{
				ctx:    ctx,
				db:     db,
				config: config.TestDefault(),

				// TODO: Add test cases.
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Genbank(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Genbank() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
