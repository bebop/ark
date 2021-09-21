package pathways

// Authors: Jordan Strasser, (Your_name_here)

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

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
	return string(stuff), err
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
func NameToId(name string) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Couldn't connect to DB: %d", err)
	}
	var id int
	query := "SELECT id FROM compound WHERE name = ?"
	err = db.Get(&id, query, name)
	if err != nil {
		return 0, err
		//whenever there's an error here we need to log desired compounds
	}
	return id, err
}

type pathdata struct {
	Rxn_id, Prod_id, Sub_id, Lvl                          int
	Type1, Prod_name, Type2, Sub_name, Name_path, Id_path string
}
type DNA struct {
	Id                         int
	Sequence, Seqhash, Genbank string
}

/*
Recursively searches throughout the database and fetches the pathways that lead to your target compound.
This query excludes any compound that occurs more than 100 times in nodes. That parameter can change,
but once we start including things like NADPH, ATP, H2O, we step into combinatorial explosions.
id_path shows you the chain of equations starting from the top, and path shows you the actual compoounds
that build up a path, which is usually just the most significant reactants and products.
*/
func GetTotalPathways(target_molecule string, levels int) ([]pathdata, error) {
	query, err := LoadSQLFile("./queries/get_total_pathways.sql")
	if err != nil {
		return nil, err
	}
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	target_id, err := NameToId(target_molecule)
	if err != nil {
		return nil, err
	}
	result := []pathdata{}
	err = db.Select(&result, query, target_id, levels)
	if err != nil {
		return result, err
	}
	db.Close()
	return result, err
}

//GetTotalPathways but limited to a single organism
func OrganismFilteredPathways(GBOrganism string, target_molecule string, levels int) ([]pathdata, error) {
	query, err := LoadSQLFile("./queries/organism_filtered_pathways.sql")
	if err != nil {
		return nil, err
	}
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	target_id, err := NameToId(target_molecule)
	if err != nil {
		return nil, err
	}
	result := []pathdata{}
	err = db.Select(&result, query, GBOrganism, target_id, levels)
	if err != nil {
		return result, err
	}
	db.Close()
	return result, err
}

/*
Input is pathway data from OrganismFilteredPathways or GetTotalPathways (if Genbank/Uniprot are complete in allbase)
and the pathway depth (levels) you want
Returns a map:
key = compound path, e.g. "XMP->guanine", which ignores the intermediary steps
value = list of DNA structs, which contains info about the DNA sequences you need
to add to an organism for it to do this chemical reaction. An individual DNA struct has Rhea reaction ID,
gene sequence, gene seqhash, and the genbank ID of the organism from which the gene comes.
*/

func GetDNA(pathways []pathdata, levels int) (map[string][]DNA, error) {
	RawQuery, err := LoadSQLFile("./queries/DNA_Gen.sql")
	if err != nil {
		return nil, err
	}
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	PathwayToDNA := make(map[string][]DNA)

	for _, pathway := range pathways {
		if pathway.Lvl == levels {
			IdList := strings.Split(pathway.Id_path, ",")
			query, args, err := sqlx.In(RawQuery, IdList)
			if err != nil {
				return PathwayToDNA, err
			}
			result := []DNA{}
			query = db.Rebind(query)
			err = db.Select(&result, query, args...)
			if err != nil {
				return PathwayToDNA, err
			}
			compounds := strings.Split(pathway.Name_path, ",")
			newString := compounds[0] + "->" + compounds[len(compounds)-1]
			PathwayToDNA[newString] = result
		}
	}
	return PathwayToDNA, err
}
