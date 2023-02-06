package vcs

import (
	"os"
	"reflect"
	"testing"
)

func TestConnectVCSDB(t *testing.T) {
	// Connect to the VCS Database
	db, err := ConnectVCSDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestConnectSynBioDB(t *testing.T) {
	// Connect to the SynBio Database
	db, err := ConnectSynBioDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestFindExistingRow(t *testing.T) {
	// Set the OS environment variable for the SynBio database
	os.Setenv("SYNBIO_DB_PATH", "./synbio.db")

	// Connect to the SynBio Database
	db, err := ConnectSynBioDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	// Create a row that should exist in the database
	row := Row{
		TableName:        "reaction",
		PrimaryKeyColumn: "ID",
		PrimaryKeyValue:  "rxn10163_c0",
	}

	// Use the function to find the row
	oldRow := findExistingRow(&row)

	// Assert if the row exists by checking if the returned row is equal
	if !reflect.DeepEqual(row, oldRow) {
		t.Error("The row does not exist")
	}

}

func TestTrackChanges(t *testing.T) {

}

func TestGetChanges(t *testing.T) {

}

func TestReadVCSState(t *testing.T) {

}

func TestWriteVCSState(t *testing.T) {

}

func TestCreateFork(t *testing.T) {

}
