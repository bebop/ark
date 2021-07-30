package main


// Authors: Jordan Strasser, David Lambert (SQL Hero)

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "log"
	_ "os"
	"strings"
)
type equ struct {
	id int
	reactants []string
	products []string
}


func main() {

	database, _ := sql.Open("sqlite3", "./allbase.db")
	get_pathways(database, "calycosin", 5)
	// run add_nodes(database) if database does not have nodes table or if it needs to be updated

}
// Recursively searches throughout the database and fetches the pathways that lead to your target compound.
// This query excludes any compound that occurs more than 100 times in nodes. That parameter can change,
// but once we start including things like NADPH, ATP, H2O, we step into combinatorial explosions.
// id_path shows you the chain of equations starting from the top, and path shows you the actual compoounds
// that build up a path, which is usually just the most significant reactants and products.
func get_pathways(database *sql.DB, compound_name string, levels int) {
	paths_query :=  `
	WITH not_one (x) AS ( SELECT DISTINCT reactant
	FROM nodes
	WHERE reactant IN (
		Select reactant From (
		SELECT count(*) as count, reactant
	FROM nodes
	GROUP BY reactant
	ORDER BY count(*) desc)
	Where count > 100)
),
	chain AS (
		SELECT id, reactant, product, 0 as lvl
	,product || ',' || reactant as path
	, cast(id as text) as id_path
	FROM nodes
	WHERE product = %q
	AND reactant NOT IN (SELECT x FROM not_one)
	UNION ALL
	SELECT n.id, n.reactant, n.product, lvl + 1 as lvl, chain.path ||
		',' || n.reactant as path
	, chain.id_path || ',' || cast(n.id as TEXT)  as  id_path
	FROM nodes n
	INNER JOIN chain ON n.product = chain.reactant
	WHERE lvl < %d
	AND n.reactant NOT IN (SELECT x FROM not_one)
	AND n.product NOT IN (SELECT x FROM not_one)
	AND INSTR(chain.path, n.reactant) = 0
	)
	select * from chain`

		tx, err := database.Begin()
		if err != nil {fmt.Println(err, "Problem loading database")}

		rows, err := tx.Query(fmt.Sprintf(paths_query, compound_name, levels))
		if err != nil { fmt.Println(err, "Something's up with the rows returned")}
		var id int
		var reactant string
		var product string
		var lvl int
		var path string
		var id_path string
		for rows.Next() {
			rows.Scan(&id, &reactant, &product, &lvl, &path, &id_path)
			fmt.Println(id, reactant, product, lvl, path, id_path)

		}

}









// takes an equation and splits it up into lists of reactants and products. Excludes bidirectional reactions
// because in Rhea, for some reason, they include both the forward and backward reactions.
func split_mols(id int, equation string) (chunk equ) {
	chunk.id  = id
	sides := strings.Split(equation, "=>")
	chunk.reactants = append(chunk.reactants, strings.Split(sides[0], " + ")...)
	chunk.products = append(chunk.products, strings.Split(sides[1], " + ")...)

	return chunk
}
// run this function if you need to add or update the nodes table
func add_nodes(database *sql.DB) {

	tx, _ := database.Begin()
	stmt, er := tx.Prepare("DROP TABLE IF EXISTS nodes")
	if er != nil {
		fmt.Println(er)
	} else {
		stmt.Exec()
	}
	statement, err := tx.Prepare("CREATE TABLE IF NOT EXISTS nodes (id INTEGER, reactant TEXT, product TEXT)")
	if err != nil {
		fmt.Println(err)
	} else {
		statement.Exec()
	}
	get_rxns := "SELECT id, equation FROM reaction WHERE directional = 1 AND trim(equation) != ''"
	rows, _ := tx.Query(get_rxns)
	var id int
	var equation string
	que := "INSERT INTO nodes (id, reactant, product) VALUES (?, ?, ?)"
	for rows.Next() {
		rows.Scan(&id, &equation)
		unit := split_mols(id, equation)
		for _, reactant := range unit.reactants {
			for _, product := range unit.products {
				statement, err = tx.Prepare(que)
				if err != nil {
					fmt.Println("could not insert: ", err)
				} else {
					statement.Exec(id, reactant, product)
				}
			}
		}
		}
		err = tx.Commit()
		if err != nil {
			fmt.Println(err)
		}
	}

