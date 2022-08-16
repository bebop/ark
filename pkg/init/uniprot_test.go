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

func TestUniprot(t *testing.T) {
	ctx := context.Background()
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	tmpConfig := config.TestDefault()
	if err != nil {
		t.Errorf("Failed to create temporary data directory for TestUniprot")
	}
	defer os.RemoveAll(tmpDataDir)

	tmpConfig.AllbasePath = filepath.Join(tmpDataDir, "uniprotTest.db")

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
			name: "TestUniprot",
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
			if err := Uniprot(tt.args.ctx, tt.args.db, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Uniprot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
