package insert

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/bebop/ark/pkg/config"
	"github.com/bebop/ark/schema"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func TestRhea(t *testing.T) {
	ctx := context.Background()
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	tmpConfig := config.DevDefault()
	if err != nil {
		t.Errorf("Failed to create temporary data directory for TestRhea")
	}
	defer os.RemoveAll(tmpDataDir)

	tmpConfig.ArkPath = filepath.Join(tmpDataDir, "rheaTest.db")

	//create test database
	err = schema.CreateDatabase(tmpConfig)

	db, err := sqlx.Open("sqlite", tmpConfig.ArkPath)
	if err != nil {
		log.Fatalf("Failed to open sqlite in ark.db: %s", err)
	}

	type args struct {
		ctx    context.Context
		db     *sqlx.DB
		config config.Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestRhea",
			args: args{
				ctx:    ctx,
				db:     db,
				config: tmpConfig,
			},
			wantErr: false,
		},
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Rhea(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Rhea() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
