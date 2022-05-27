CREATE TABLE IF NOT EXISTS seqhash (seqhash TEXT NOT NULL PRIMARY KEY, sequence TEXT NOT NULL, circular BOOL NOT NULL DEFAULT FALSE, doublestranded BOOL NOT NULL DEFAULT TRUE, seqhashtype TEXT NOT NULL CHECK (seqhashtype IN ('DNA', 'RNA', 'PROTEIN')), translation TEXT REFERENCES seqhash(seqhash));

CREATE TABLE IF NOT EXISTS genbank (accession TEXT PRIMARY KEY, seqhash TEXT NOT NULL REFERENCES seqhash(seqhash));

CREATE TABLE IF NOT EXISTS genbank_features (seqhash TEXT NOT NULL REFERENCES seqhash(seqhash), parent TEXT NOT NULL REFERENCES genbank(accession), PRIMARY KEY( seqhash ,  parent ));

CREATE TABLE IF NOT EXISTS uniprot (accession TEXT PRIMARY KEY, database TEXT NOT NULL, seqhash TEXT NOT NULL REFERENCES seqhash(seqhash));

CREATE TABLE IF NOT EXISTS chebi (accession TEXT PRIMARY KEY, subclass_of TEXT REFERENCES chebi(accession));

CREATE TABLE IF NOT EXISTS compound (id INT NOT NULL, accession TEXT PRIMARY KEY, position TEXT, name TEXT, html_name TEXT, formula TEXT, charge TEXT, chebi TEXT REFERENCES chebi(accession), polymerization_index TEXT, compound_type TEXT NOT NULL CHECK(compound_type IN ('small_molecule', 'polymer', 'generic_polypeptide', 'generic_polynucleotide', 'generic_heteropolysaccharide')));

CREATE TABLE IF NOT EXISTS reactive_part (id INT, accession TEXT PRIMARY KEY, name TEXT, html_name TEXT, compound TEXT NOT NULL REFERENCES compound(accession));

CREATE TABLE IF NOT EXISTS reaction (id INT, directional BOOL NOT NULL DEFAULT FALSE, accession TEXT PRIMARY KEY, status TEXT, comment TEXT, equation TEXT, html_equation TEXT, is_chemically_balanced BOOL NOT NULL DEFAULT TRUE, is_transport BOOL NOT NULL DEFAULT FALSE, ec TEXT, location TEXT);

CREATE TABLE IF NOT EXISTS reactionside (accession TEXT PRIMARY KEY);

CREATE TABLE IF NOT EXISTS reactionside_reaction (reaction TEXT NOT NULL REFERENCES reaction(accession), reactionside TEXT NOT NULL REFERENCES reactionside(accession), reactionside_reaction_type TEXT NOT NULL CHECK(reactionside_reaction_type IN ('substrate_or_product', 'substrate', 'product')), PRIMARY KEY( reaction ,  reactionside ));

CREATE TABLE IF NOT EXISTS reaction_participant (compound TEXT REFERENCES compound(accession), reactionside TEXT NOT NULL REFERENCES reactionside(accession), contains INT, contains_n BOOL NOT NULL DEFAULT FALSE, minus BOOL NOT NULL DEFAULT FALSE, plus BOOL NOT NULL DEFAULT FALSE, PRIMARY KEY( compound ,  reactionside ));

CREATE TABLE IF NOT EXISTS uniprot_to_reaction (reaction TEXT REFERENCES reaction(accession), uniprot TEXT REFERENCES uniprot(accession), PRIMARY KEY( reaction ,  uniprot ))