package init

import (
	"context"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/pkg/db"
	"github.com/TimothyStiles/surrealdb.go"
)

func TestUniprot(t *testing.T) {
	ctx := context.Background()
	testConfig := config.TestDefault()
	testDB, err := db.CreateTestDB("uniprot", testConfig)
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
			name: "TestUniprot",
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
			if err := Uniprot(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Uniprot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
