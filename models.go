package allbase

import (
	"database/sql"
	"github.com/jmoiron/sqlx/types"
)

var schema = `

-- Create Seqhash Table --
CREATE TABLE seqhash (
	seqhash TEXT PRIMARY KEY,
	sequence TEXT NOT NULL,
	circular BOOLEAN NOT NULL DEFAULT FALSE,
	doublestranded BOOLEAN NOT NULL DEFAULT TRUE,
	seqhashtype TEXT NOT NULL CHECK (seqhashtype IN ('DNA', 'RNA', 'PROTEIN')),
	translation TEXT REFERENCES seqhash(seqhash)
);

-- Create Genbank Table --
CREATE TABLE genbank (
	accession TEXT PRIMARY KEY,
	genbankhash TEXT NOT NULL, -- adler32 checksum
	genbank TEXT NOT NULL, 
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash)
);

-- Create Genbank Features Table --
CREATE TABLE genbankfeatures (
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash),
	genbank TEXT NOT NULL REFERENCES genbank(id)
);

-- Create Uniprot Table --
CREATE TABLE uniprot (
	accession TEXT PRIMARY KEY,
	uniprothash TEXT NOT NULL, -- adler32 checksum
	uniprot JSON NOT NULL,
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash)
);

-- Rhea

CREATE TABLE IF NOT EXISTS chebi (
        accession TEXT PRIMARY KEY,
        subclassof TEXT REFERENCES chebi(accession)
);

CREATE TABLE IF NOT EXISTS compound (
        id INT,
        accession TEXT PRIMARY KEY,
        position TEXT,
        name TEXT,
        htmlname TEXT,
        formula TEXT,
        charge TEXT,
        chebi TEXT REFERENCES chebi(accession),
        polymerizationindex TEXT,
        compoundtype TEXT NOT NULL CHECK(compoundtype IN ('SmallMolecule', 'Polymer', 'GenericPolypeptide', 'GenericPolynucleotide', 'GenericHeteropolysaccharide'))
);

CREATE TABLE IF NOT EXISTS reactivepart (
        id INT,
        accession TEXT PRIMARY KEY,
        name TEXT,
        htmlname TEXT,
        compound TEXT NOT NULL REFERENCES compound(accession)
);

CREATE TABLE IF NOT EXISTS reaction (
        id INT,
        directional BOOL NOT NULL DEFAULT false,
        accession TEXT PRIMARY KEY,
        status TEXT,
        comment TEXT,
        equation TEXT,
        htmlequation TEXT,
        ischemicallybalanced BOOL NOT NULL DEFAULT true,
        istransport BOOL NOT NULL DEFAULT false,
        ec TEXT,
        location TEXT
);

CREATE TABLE IF NOT EXISTS reactionside (
        accession TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS reactionsidereaction (
        reaction TEXT NOT NULL REFERENCES reaction(accession),
        reactionside TEXT NOT NULL REFERENCES reactionside(accession),
        reactionsidereactiontype TEXT NOT NULL CHECK(reactionsidereactiontype IN ('substrateorproduct', 'substrate', 'product'))
);

CREATE TABLE IF NOT EXISTS reactionparticipant (
        compound TEXT REFERENCES compound(accession),
        reactionside TEXT NOT NULL REFERENCES reactionside(accession),
        contains INTEGER,
        containsn BOOL NOT NULL DEFAULT false,
        minus BOOL NOT NULL DEFAULT false,
        plus BOOL NOT NULL DEFAULT false
);

-- Uniprot to reaction

CREATE TABLE IF NOT EXISTS uniprot_to_reaction (
        reaction TEXT REFERENCES reaction(accession),
        uniprot TEXT REFERENCES uniprot(accession)
);
`

/******************************************************************************

Schema structs begin here

******************************************************************************/

type Seqhash struct {
	Seqhash        string         `db:"seqhash"`
	Circular       bool           `db:"circular"`
	DoubleStranded bool           `db:"doublestranded"`
	Sequence       string         `db:"sequence"`
	SeqhashType    string         `db:"seqhashtype"`
	Translation    sql.NullString `db:"translation"`
}

type Genbank struct {
	Id          string `db:"id"`
	Genbank     string `db:"genbank"`
	GenbankHash string `db:"genbankhash"`
	PolyVersion string `db:"polyversion"`
	Seqhash     string `db:"seqhash"`
}

type GenbankFeature struct {
	Seqhash string `db:"seqhash"`
	Genbank string `db:"genbank"`
}

type Uniprot struct {
	Accession   string         `db:"accession"`
	UniprotHash string         `db:"uniprothash"`
	Uniprot     types.JSONText `db:"uniprot"`
	Seqhash     string         `db:"seqhash"`
}

type GenbankFull struct {
	Seqhashes       []Seqhash
	Genbank         Genbank
	GenbankFeatures []GenbankFeature
}
