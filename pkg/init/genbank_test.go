package init

import (
	"context"
	"os"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/db"
	"github.com/uptrace/bun"
)

func TestGenbank(t *testing.T) {
	ctx := context.Background()

	testDB, err := db.CreateTestDB("genbank.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDB.DirPath)

	err = CreateGenbankTable(ctx, testDB.DB)
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
				db:     testDB.DB,
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
