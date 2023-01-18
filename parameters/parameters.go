package parameters

import (
	"os"
	"path/filepath"
)

const DefaultServerPort = "8080"

// Get the file path of the Retsynth database from the environment variable, if it exists otherwise set default path
func RetsynthDBPath() string {

	// Get the path from the environment variable
	databasePath, ok := os.LookupEnv("RETSYNTH_DB_PATH")
	if !ok {
		databasePath = "../data/dev/retsynth/minimal.db"
	}

	// Convert to absolute path
	absolutepath, err := filepath.Abs(databasePath)
	if err != nil {
		panic(err)
	}

	return absolutepath
}
