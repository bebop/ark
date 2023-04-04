-- Meta: this code takes about 30 seconds to finish, irrespective of the levels depth parameter. 
-- When the organism filter is put on the normal pathway generator, speed is reduced by 1/3...
-- native_rxns CTE fetches rxn ids that are native to a user-specified organism
WITH native_rxns as (
select distinct case
when r.directional = 0
then r.id + 1
else r.id
end id
 from uniprot u 
inner join uniprot_to_reaction ur on u.accession = ur.uniprot 
inner join reaction r on ur.reaction = r.accession
where u.seqhash IN (
select s.translation  
from genbank g inner join genbankfeatures gf 
on g.accession = gf.genbank 
inner join seqhash s on gf.seqhash = s.seqhash 
WHERE g.accession = ? -- organism param
) 
),
-- native_rxns selects from 'case' because uniprot_to_reaction includes the base parent reactions that are undirectional. 
-- sometimes only these parent reaction have an associated enzyme. 
-- modified to r.id + 1 so that it can match items within the stitch cte 
 
 -- formats the ark tables in a conenient way to see reaction participants, ids, names
 stitch as ( 
 select a.accession, a.id as rxn_id,
		b.reactionside,
		b.reactionsidereactiontype,
		c.compound,
		d.id as cmpd_id,
		d.name as name
		FROM reaction a
		inner join reactionsidereaction b on a.accession = b.reaction
		inner join reactionparticipant c USING(reactionside)
		inner join compound d ON c.compound = d.accession
		where b.reactionsidereactiontype <> 'substrateorproduct'
        and d.id != 0 
		and rxn_id IN (select id from native_rxns) -- organism filter
		or rxn_id IN (select id+1 from native_rxns) -- gets reverse reactions also
		),

		-- finds most common compounds in an organism to avoid trivial pathways 
		common_compounds AS ( 
		select cmpd_id FROM ( select count(*) as count, cmpd_id, name 
		from stitch 
		group by cmpd_id 
		order by count(*) desc
		)
	    where count >  30 -- param to filter common substrates
		),
		-- recursively goes through stitch, finding the substrates of substrates, until the level param is reached 
		-- filters reactions to only return ones native to the organism of interest 
		chain as ( 
		select c.rxn_id as rxn_id, c.reactionsidereactiontype as type1, c.cmpd_id as prod_id, c.name as prod_name,
			d.reactionsidereactiontype as type2, d.cmpd_id as sub_id, d.name as sub_name, 0 as lvl,
			c.name|| ',' ||d.name as name_path,
			cast(c.rxn_id as text) as id_path
			from
			stitch c
			inner join stitch d on c.rxn_id = d.rxn_id
			where c.cmpd_id = ? 
			and d.cmpd_id NOT IN (select * from common_compounds)
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
				where lvl < ? 
				and f.reactionsidereactiontype LIKE '%substrate%'
				and e.reactionsidereactiontype NOT LIKE '%substrate%'
				and f.cmpd_id NOT IN (select * from common_compounds)
				and instr(chain.name_path, f.name) = 0
				)
				select * from chain