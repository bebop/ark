package init

import (
	"context"
	"os"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/db"
	"github.com/uptrace/bun"
)

func TestUniprot(t *testing.T) {
	ctx := context.Background()
	testDB, err := db.CreateTestDB("uniprot.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDB.DirPath)

	err = CreateUniprotTable(ctx, testDB.DB)
	if err != nil {
		t.Fatal(err)
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
			name: "TestUniprot",
			args: args{
				ctx:    ctx,
				db:     testDB.DB,
				config: config.TestDefault(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Uniprot(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Uniprot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
