package pathways

import (
	"testing"
)

func TestGetTotalPathways(t *testing.T) {
	result := GetTotalPathways("calycosin", 4)
	if len(result) != 9 {
		t.Error("Expected: len(result) = 9, Got: ", len(result))
	}
}

func TestNameToId(t *testing.T) {
	if id := NameToId("XMP"); id != 5036 {
		t.Error("Expected: 5036, Got: ", id)
	}
}

func TestOrganismFilteredPathways(t *testing.T) {
	result := OrganismFilteredPathways("CP060121", "XMP", 1)
	if len(result) != 37 {
		t.Error("Expected 37 path branches, Got ", len(result))
	}
}
