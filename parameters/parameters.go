package parameters

import "os"

const DefaultServerPort = "8080"

// Get the file path of the Retsynth database from the environment variable, if it exists otherwise set default path
func RetsynthDBPath() string {
	databasePath, ok := os.LookupEnv("RETSYNTH_DB_PATH")
	if !ok {
		databasePath = "../data/dev/retsynth/minimal.db"
	}
	return databasePath
}
