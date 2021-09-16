select distinct r.accession, s.sequence, gf.seqhash, gf.genbank
from reaction r
inner join uniprot_to_reaction ur ON ur.reaction = r.accession
inner join uniprot u ON ur.uniprot = u.accession
inner join seqhash s on s.translation = u.seqhash
inner join genbankfeatures gf on gf.seqhash = s.seqhash
WHERE r.accession IN 
(
Select accession from reaction where id = ?
)