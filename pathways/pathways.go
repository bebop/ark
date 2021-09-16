package pathways

// Authors: Jordan Strasser, David Lambert (SQL Assistance)

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

//Reads and loads SQL files as string
func LoadSQLFile(path string) (string, error) {
	realpath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	stuff, err := ioutil.ReadFile(realpath)
	if err != nil {
		return "", err
	}
	return string(stuff), nil
}

//Easy database connector
func ConnectDB() (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	db, err = sqlx.Connect("sqlite3", "./allbasetest.db")
	if err != nil {
		return db, err
	}
	return db, err
}

//Gets id from compound name if it exists in allbase
func NameToId(name string) int {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Couldn't connect to DB: %d", err)
	}
	var id int
	query := "SELECT id FROM compound WHERE name = ?"
	err = db.Get(&id, query, name)
	if err != nil {
		log.Fatalf("Query didn't work, could be SQL query, could be Golang struct: %d", err)
		//whenever there's an error here we need to log desired compounds
	}
	return id
}

type pathdata struct {
	Rxn_id, Prod_id, Sub_id, Lvl                          int
	Type1, Prod_name, Type2, Sub_name, Name_path, Id_path string
}

/*
Recursively searches throughout the database and fetches the pathways that lead to your target compound.
This query excludes any compound that occurs more than 100 times in nodes. That parameter can change,
but once we start including things like NADPH, ATP, H2O, we step into combinatorial explosions.
id_path shows you the chain of equations starting from the top, and path shows you the actual compoounds
that build up a path, which is usually just the most significant reactants and products.
*/
func GetTotalPathways(target_molecule string, levels int) []pathdata {
	query, err := LoadSQLFile("./queries/get_total_pathways.sql")
	if err != nil {
		log.Fatalf("Could not load SQL file: %d", err)
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %d", err)
	}
	target_id := NameToId(target_molecule)
	result := []pathdata{}
	err = db.Select(&result, query, target_id, levels)
	if err != nil {
		log.Fatalf("Could not perform query either because SQL error or Golang struct: %d", err)
	}
	db.Close()
	return result
}

//GetTotalPathways but limited to a single organism
func OrganismFilteredPathways(GBOrganism string, target_molecule string, levels int) []pathdata {
	query, err := LoadSQLFile("./queries/organism_filtered_pathways.sql")
	if err != nil {
		log.Fatalf("Could not load SQL file: %d", err)
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %d", err)
	}
	target_id := NameToId(target_molecule)
	result := []pathdata{}
	err = db.Select(&result, query, GBOrganism, target_id, levels)
	if err != nil {
		log.Fatalf("Could not perform query either because SQL error or Golang struct: %d", err)
	}
	db.Close()
	return result
}
