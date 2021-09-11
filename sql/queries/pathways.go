package pathways

// Authors: Jordan Strasser, David Lambert (SQL Assistance)

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func pathways() {

	database, err := sql.Open("sqlite3", "./allbase.db")
	if err != nil {
		fmt.Println("Issue opening the database")
	}
	if len(os.Args) < 5 {
		fmt.Println("Expected 2 arguments: -compound and -levels")
		os.Exit(1)
	}

	makeCmd := flag.NewFlagSet("make", flag.ExitOnError)
	target_compound := makeCmd.String("compound", "calycosin", "ID of target compound from allbase")
	levels := makeCmd.Int("levels", 1, "Number of pathway steps max limit")
	makeCmd.Parse(os.Args[2:])
	fmt.Println(target_compound, *target_compound)
	fmt.Println(levels, *levels)

	id_query := "SELECT id FROM compound WHERE name = " + "'" + *target_compound + "'"

	rows, err := database.Query(id_query)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	var id, final_id int
	for rows.Next() {
		count = count + 1
		rows.Scan(&id)
		final_id = id
	}
	if count == 0 {
		fmt.Println("Compound could not be found in local database. We'll be improving our alchemy shortly.")
	}

	//command line arguments are not being processed

	// target_id and levels should be modifiable
	get_total_pathways(database, final_id, *levels)

}

// recursively searches throughout the database and fetches the pathways that lead to your target compound.
// This query excludes any compound that occurs more than 100 times in nodes. That parameter can change,
// but once we start including things like NADPH, ATP, H2O, we step into combinatorial explosions.
// id_path shows you the chain of equations starting from the top, and path shows you the actual compoounds
// that build up a path, which is usually just the most significant reactants and products.
func get_total_pathways(database *sql.DB, target_id int, levels int) {
	paths_query := `
	WITH stitch AS (
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
		and d.id != 0
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
			-- and c.rxn_id IN (SELECT id FROM native_rxns)
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
				and f.reactionsidereactiontype LIKE '%substrate%'
				and e.reactionsidereactiontype NOT LIKE '%substrate%'
				and f.cmpd_id NOT IN (SELECT ids FROM not_one)
				-- and e.rxn_id IN (SELECT id FROM native_rxns)
				and instr(chain.name_path, f.name) = 0
				)
				select * from chain`
	tx, err := database.Begin()
	if err != nil {
		fmt.Println(err, "Problem loading database")
	}

	rows, err := tx.Query(fmt.Sprintf(paths_query, target_id, levels))
	if err != nil {
		fmt.Println(err, "Something's up with the rows returned")
	}
	var rxn_id, prod_id, sub_id, lvl int
	var type1, prod_name, type2, sub_name, name_path, id_path string
	for rows.Next() {
		rows.Scan(&rxn_id, &type1, &prod_id, &prod_name, &type2, &sub_id, &sub_name, &lvl, &name_path, &id_path)
		fmt.Println(rxn_id, type1, prod_id, prod_name, type2, sub_id, sub_name, lvl, name_path, id_path)

	}

}
func organism_filtered_pathways(database *sql.DB, target_id int, levels int) {
	paths_query := `
	WITH native_rxns as (
select  distinct r.id
from genbank g inner join genbankfeatures gf 
on g.accession = gf.genbank 
inner join seqhash s on gf.seqhash = s.seqhash 
inner join seqhash j on s.translation = j.seqhash
inner join  uniprot u on j.seqhash = u.seqhash -- this step adds 10 seconds to the computation
inner join uniprot_to_reaction ur on u.accession = ur.uniprot
inner join reaction r on r.accession = ur.reaction
-- issue here is that parent reactions are included where id%4=0.
WHERE g.accession LIKE 'CP060121' -- can do any org accession or multiple orgs, but remember '%'
),
good as (
select id from native_rxns
where id % 4 != 0
),
better as ( 
select id+1 as fwd from native_rxns
where id % 4 = 0
union ALL 
select id from good
),
allowit as ( 
select id + 2 as ids from native_rxns
where id % 4 = 0
union ALL
select * from better
),
 stitch AS (
	select a.accession, a.id as rxn_id,
		b.reactionside,
		b.reactionsidereactiontype,
		c.compound,
		d.id as cmpd_id,
		d.name
		FROM reaction a
	    inner join allowit n on n.ids = a.id -- do this to limit to an org
		inner join reactionsidereaction b on a.accession = b.reaction
		inner join reactionparticipant c USING(reactionside)
		inner join compound d ON c.compound = d.accession
		where b.reactionsidereactiontype <> 'substrateorproduct'
		and d.id != 0
		),
		chain as ( 
		select c.rxn_id as rxn, c.reactionsidereactiontype as type1, c.cmpd_id as prod_id, c.name as prod_name,
			d.reactionsidereactiontype as type2, d.cmpd_id as sub_id, d.name as sub_name, 0 as lvl,
			c.name|| ',' ||d.name as name_path,
			cast(c.rxn_id as text) as id_path
			from
			stitch c
			inner join stitch d on c.rxn_id = d.rxn_id
			where c.cmpd_id = 1663 -- param
			and sub_name NOT IN ('H2O', 'H(+)', 'CO2', 'NAD(+)', 'NADH', 'ATP', 'FADH2', 'ADP', 'O2')
			and type2  = 'substrate'
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
				where lvl < 2 -- param
				and f.reactionsidereactiontype LIKE '%substrate%'
				and e.reactionsidereactiontype NOT LIKE '%substrate%'
				and f.name NOT IN ('H2O', 'H(+)', 'CO2', 'NAD(+)', 'NADH', 'ATP', 'FADH2', 'ADP')

				and instr(chain.name_path, f.name) = 0
				)
				select * from chain`
	tx, err := database.Begin()
	if err != nil {
		fmt.Println(err, "Problem loading database")
	}

	rows, err := tx.Query(fmt.Sprintf(paths_query, target_id, levels))
	if err != nil {
		fmt.Println(err, "Something's up with the rows returned")
	}
	var rxn_id, prod_id, sub_id, lvl int
	var type1, prod_name, type2, sub_name, name_path, id_path string
	for rows.Next() {
		rows.Scan(&rxn_id, &type1, &prod_id, &prod_name, &type2, &sub_id, &sub_name, &lvl, &name_path, &id_path)
		fmt.Println(rxn_id, type1, prod_id, prod_name, type2, sub_id, sub_name, lvl, name_path, id_path)

	}

}
