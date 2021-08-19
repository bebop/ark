package main

import "database/sql"

/* Authors: Jordan Strasser, Keoni Gandall, David Lambert (SQL Assist)
*/ 

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_ "log"
	_ "os"
)

func main() {

	database, err := sql.Open("sqlite3", "./allbase.db")
	if err != nil {
		fmt.Println("Issue opening the database")
	}
	// target_id and levels should be modifiable
	get_pathways(database, 10019, 5)


}
// recursively searches throughout the database and fetches the pathways that lead to your target compound.
// This query excludes any compound that occurs more than 100 times in nodes. That parameter can change,
// but once we start including things like NADPH, ATP, H2O, we step into combinatorial explosions.
// id_path shows you the chain of equations starting from the top, and path shows you the actual compoounds
// that build up a path, which is usually just the most significant reactants and products.
func get_pathways(database *sql.DB, target_id int, levels int) {
	paths_query :=  `WITH stitch AS ( 
select a.accession, a.id as rxn_id,
b.reactionside, 
b.reactionsidereactiontype,
c.compound,
d.id as cmpd_id,
d.name
FROM reaction a
inner join reactionsidereaction b on a.accession = b.reaction 
inner join reactionparticipant c USING(reactionside)
inner join compound d ON c.compound = d.accession
where b.reactionsidereactiontype <> 'substrateorproduct'
), 
not_one(ids) AS (
SELECT cmpd_id FROM ( 
SELECT count(*) as count, cmpd_id
	FROM stitch
	GROUP BY cmpd_id
	ORDER BY count(*) DESC )
	WHERE count > 100 -- modify to filter out commonly available molecules
	),
	chain AS (
select c.rxn_id, c.reactionsidereactiontype as type1, c.cmpd_id as prod_id, c.name as prod_name,
d.reactionsidereactiontype as type2, d.cmpd_id as sub_id, d.name as sub_name, 0 as lvl, 
c.name|| ',' ||d.name as name_path, 
cast(c.rxn_id as text) as id_path
from 
stitch c 
inner join stitch d on c.rxn_id = d.rxn_id 
where c.cmpd_id = %d
and sub_id NOT IN (SELECT ids from not_one)
and type2 = 'substrate'
and type1 <> 'substrate'
UNION ALL 
select e.rxn_id, e.reactionsidereactiontype as type1, e.cmpd_id as prod_id, e.name as prod_name,
f.reactionsidereactiontype as type2, f.cmpd_id as sub_id, f.name as sub_name, chain.lvl + 1 as lvl,
chain.name_path || ',' || f.name as name_path, 
chain.id_path || ',' || cast(e.rxn_id as text) as id_path 
from 
stitch e
inner join stitch f on e.rxn_id = f.rxn_id 
inner join chain on e.cmpd_id = chain.sub_id
where lvl < %d
and f.reactionsidereactiontype = 'substrate'
and e.reactionsidereactiontype <> 'substrate'
and f.cmpd_id NOT IN (SELECT ids FROM not_one)
and instr(chain.name_path, f.name) = 0
) 
select * from chain`

		tx, err := database.Begin()
		if err != nil {fmt.Println(err, "Problem loading database")}

		rows, err := tx.Query(fmt.Sprintf(paths_query, target_id, levels))
		if err != nil { fmt.Println(err, "Something's up with the rows returned")}
		var rxn_id int
		var type1 string
		var prod_id int
		var prod_name string
		var type2 string
		var sub_id int
		var sub_name string
		var lvl int
		var name_path string
		var id_path string
		for rows.Next() {
			rows.Scan(&rxn_id, &type1, &prod_id, &prod_name, &type2, &sub_id, &sub_name, &lvl, &name_path, &id_path)
			fmt.Println(rxn_id, type1, prod_id, prod_name, type2, sub_id, sub_name, lvl, name_path, id_path)

		}

}
