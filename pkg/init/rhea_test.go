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

func TestRhea(t *testing.T) {
	ctx := context.Background()
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	err = CreateRheaTable(ctx, db)
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
		// TODO: Add test cases.
		{
			name: "TestRhea",
			args: args{
				ctx:    ctx,
				db:     db,
				config: config.TestDefault(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Rhea(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Rhea() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
