package download_test

import (
	"testing"

	"github.com/allyourbasepair/allbase/pkg/download"
)

func TestGetPageLinks(t *testing.T) {
	links, err := download.Links("http://example.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(links) == 0 {
		t.Errorf("Error: No links found")
	}
}
