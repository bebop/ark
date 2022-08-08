package init

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/allbase/schema"
	"github.com/jmoiron/sqlx"
)

func TestGenbank(t *testing.T) {
	ctx := context.Background()
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	tmpConfig := config.TestDefault()
	if err != nil {
		t.Errorf("Failed to create temporary data directory for TestGenbank")
	}
	defer os.RemoveAll(tmpDataDir)

	tmpConfig.AllbasePath = filepath.Join(tmpDataDir, "genbankTest.db")

	//create test database
	err = schema.CreateDatabase(tmpConfig)

	db, err := sqlx.Open("sqlite", tmpConfig.AllbasePath)
	if err != nil {
		log.Fatalf("Failed to open sqlite in allbase.db: %s", err)
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
			name: "TestGenbank",
			args: args{
				ctx:    ctx,
				db:     db,
				config: tmpConfig,
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
