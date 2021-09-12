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
				select * from chain