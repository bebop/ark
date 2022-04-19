package pathways

import (
	"testing"
)

func TestLoadSQLFile(t *testing.T) {
	//this file exists
	_, err := LoadSQLFile("./queries/get_total_pathways.sql")
	if err != nil {
		t.Error(err)
	}
	//this file is intentionally not there, should throw errors
	_, err = LoadSQLFile("./not_there.sql")
	if err == nil {
		t.Error("Error check not working")
	}
}

func TestGetTotalPathways(t *testing.T) {
	result, err := GetTotalPathways("calycosin", 4)
	if len(result) != 9 {
		t.Error("Expected: len(result) = 9, got ", len(result), err)
	}
	_, err = GetTotalPathways("NotAvail", 4)
	if err == nil {
		t.Error("Error check not working for unavail compounds")
	}
}
func TestNameToId(t *testing.T) {
	if id, err := NameToId("XMP"); id != 5036 {
		t.Error("Expected: 5036, got: ", err)
	}
	_, err := NameToId("NotThere")
	if err == nil {
		t.Error("Error checker is not working for NameToID")
	}
}
func TestOrganismFilteredPathways(t *testing.T) {
	result, err := OrganismFilteredPathways("CP060121", "XMP", 1)
	if len(result) != 37 {
		t.Error("Expected 37 path branches, got: ", len(result), err)
	}
	//intentionally broken. having GBOrganism not there doesn't actually throw an error, just an empty list
	result, _ = OrganismFilteredPathways("NotThere", "XMP", 1)
	if len(result) != 0 {
		t.Error("Error check failed for broken GB Organism ID")
	}
	//intentionally broken
	_, err = OrganismFilteredPathways("CP060121", "NotThere", 1)
	if err == nil {
		t.Error("Error check failed for broken target molecule")
	}
}
func TestGetDNA(t *testing.T) {
	pathways, _ := OrganismFilteredPathways("CP060121", "XMP", 1)
	DNAseqs, err := GetDNA(pathways, 1)
	if len(DNAseqs) != 25 {
		t.Error("Expected 25 key-value pairs, got: ", len(DNAseqs), err)
	}
}
