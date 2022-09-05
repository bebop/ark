package db

import (
	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/TimothyStiles/surrealdb.go"
)

// CreateTestDB creates a temporary database for testing.
func CreateTestDB(namespace string, testConfig config.Config) (*surrealdb.DB, error) {
	// create a temporary database
	testDB, err := surrealdb.New(testConfig.AllbaseURL)

	if err != nil {
		return nil, err
	}

	_, err = testDB.Signin(map[string]interface{}{
		"user": testConfig.AdminUser,
		"pass": testConfig.AdminPassword,
	})

	if err != nil {
		return nil, err
	}

	// use the temporary database
	_, err = testDB.Use(namespace, testConfig.DBName)

	return testDB, nil
}
