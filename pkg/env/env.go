package env

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// RootPath returns the root directory of the project as an absolute path.
func RootPath() string {
	_, runfile, _, _ := runtime.Caller(0)
	dir := filepath.Join("/", path.Dir(runfile))

	splitPath := strings.Split(dir, "/")
	var lastarkIndex int

	// get last slice element equal to "ark" of splitPath
	for index, substring := range splitPath {
		if substring == "ark" {
			lastarkIndex = index
		}
	}

	splitRootPath := splitPath[:lastarkIndex+1]
	rootPathWithoutSlash := filepath.Join(splitRootPath...)
	rootPath := path.Join("/", rootPathWithoutSlash)

	return rootPath
}
