package db

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/TimothyStiles/allbase/pkg/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

// CreateTestDB creates a temporary database for testing.
func CreateTestDB(dbName string) (TestDB, error) {
	testDB := TestDB{}
	tmpDataPath := filepath.Join(env.RootPath(), "data")
	var err error

	testDB.DirPath, err = ioutil.TempDir(tmpDataPath, "tmp-*")
	if err != nil {
		return TestDB{}, err
	}

	testDB.DBPath = filepath.Join(testDB.DirPath, dbName)

	sqldb, err := sql.Open(sqliteshim.ShimName, testDB.DBPath)

	if err != nil {
		panic(err)
	}

	testDB.DB = bun.NewDB(sqldb, sqlitedialect.New())
	return testDB, nil
}

// TestDB is a temporary database for testing.
type TestDB struct {
	DB      *bun.DB
	DBPath  string
	DirPath string
}
