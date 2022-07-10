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
	dir := path.Dir(runfile)
	rootPath := "/"
	for _, subString := range strings.Split(dir, "/") {
		rootPath = filepath.Join(rootPath, subString)
		if subString == "allbase" {
			break
		}
	}
	return rootPath
}
