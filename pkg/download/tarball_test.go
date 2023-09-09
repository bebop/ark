package download_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/bebop/ark/pkg/download"
)

func TestTarball(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.RemoveAll(tmpDataDir)

	response, err := http.Get("https://github.com/TimothyStiles/poly/archive/refs/tags/v0.0.0.tar.gz")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer response.Body.Close()

	err = download.Tarball(response.Body, "README", tmpDataDir)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
