
-- need to do this step because of a technicality in the database's schema.
-- ids involved in reaction pathways may not directly map to uniprot_to_reaction,
-- but they will be associated with a "parent reaction" that IS associated. this generates that table to show potential parents
WITH potentialparents AS (
select r.id as id1, r.accession as accession1, j.id as id2, j.accession as accession2, k.id as id3, k.accession as accession3
 from reaction r 
 inner join reaction j on r.id-1 = j.id
 inner join reaction k on j.id-1 = k.id
where r.id IN (?)
)

-- formats the output so that specific reaction accessions show the associated DNA. 
-- NOTE: these reactions may be different by -1 or -2 than the input reaction, 
-- indicating that the parent reaction has been found and 
-- that there is no documented directionality to the reaction within Rhea
select distinct r.id, s.sequence, gf.seqhash, gf.genbank
from reaction r
inner join uniprot_to_reaction ur ON ur.reaction = r.accession
inner join uniprot u ON ur.uniprot = u.accession
inner join seqhash s on s.translation = u.seqhash
inner join genbankfeatures gf on gf.seqhash = s.seqhash
WHERE r.accession IN 
(
-- this checks which actual accessions number to use for the reaction you want
select CASE
when accession1 IN (select reaction from uniprot_to_reaction) 
then accession1
when accession2 IN (select reaction from uniprot_to_reaction)
then accession2
when accession3 IN (select reaction from uniprot_to_reaction)
then accession3
END parent_accession
from potentialparents
)