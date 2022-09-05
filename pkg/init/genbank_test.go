package init

import (
	"context"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/db"
	"github.com/TimothyStiles/surrealdb.go"
)

func TestGenbank(t *testing.T) {
	ctx := context.Background()
	testConfig := config.TestDefault()
	testDB, err := db.CreateTestDB("genbank", testConfig)
	if err != nil {
		t.Errorf("error creating test database: %v", err)
	}

	type args struct {
		ctx    context.Context
		db     *surrealdb.DB
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
				db:     testDB,
				config: testConfig,

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
