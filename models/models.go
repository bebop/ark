package models

import (
	"compress/gzip"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/TimothyStiles/poly"
	"github.com/TimothyStiles/poly/io/uniprot"
	"github.com/TimothyStiles/poly/seqhash"
	"github.com/allyourbasepair/ark/pkg/rhea"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

var Schema = `
PRAGMA journal_mode=WAL;
PRAGMA foreign_keys = ON;

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
	-- genbankhash TEXT NOT NULL, -- adler32 checksum
	-- genbank TEXT NOT NULL, 
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash)
);

-- Create Genbank Features Table --
CREATE TABLE genbankfeatures (
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash),
	genbank TEXT NOT NULL REFERENCES genbank(accession)
);

-- Create Uniprot Table --
CREATE TABLE uniprot (
	accession TEXT PRIMARY KEY,
	database TEXT NOT NULL,
	-- uniprothash TEXT NOT NULL, -- adler32 checksum
	-- uniprot JSON NOT NULL,
	seqhash TEXT NOT NULL REFERENCES seqhash(seqhash)
);

-- Rhea

CREATE TABLE IF NOT EXISTS chebi (
        accession TEXT PRIMARY KEY,
        subclassof TEXT REFERENCES chebi(accession)
);

CREATE TABLE IF NOT EXISTS compound (
        id INT NOT NULL,
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

/******************************************************************************

Rhea

******************************************************************************/

// RheaInsert inserts the Rhea database into an SQLite database with proper normalization for advanced queries.
func RheaInsert(db *sqlx.DB, rhea rhea.Rhea) error {
	// Start transaction with database for insertion. This ensures if there are any problems, they are seamlessly rolled back
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// First, insert ChEBIs and Compounds
	compoundKeys := make(map[string]bool)
	for _, compound := range rhea.Compounds {
		// Insert root chebi. Ie, what this current compound's subclass is
		_, err = tx.Exec("INSERT OR IGNORE INTO chebi(accession) VALUES (?)", compound.SubclassOfChEBI)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// Insert the chebi of the current compound. If it is already inserted, update the subclassification
		_, err = tx.Exec("INSERT INTO chebi(accession, subclassof) VALUES (?, ?) ON CONFLICT (accession) DO UPDATE SET subclassof = ?", compound.ChEBI, compound.SubclassOfChEBI, compound.SubclassOfChEBI)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// Insert the compound itself
		_, err = tx.Exec("INSERT INTO compound(id, accession, position, name, htmlname, formula, charge, chebi, polymerizationindex, compoundtype) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING", compound.ID, compound.Accession, compound.Position, compound.Name, compound.HTMLName, compound.Formula, compound.Charge, compound.ChEBI, compound.PolymerizationIndex, compound.CompoundType)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// If the compound isn't a small molecule or polymer, that means it would be a reactive part of a larger compound. So we add it in
		if (compound.CompoundType != "SmallMolecule") && (compound.CompoundType != "Polymer") {
			_, err = tx.Exec("INSERT INTO reactivepart(id, accession, name, htmlname, compound) VALUES (?, ?, ?, ?, ?)", compound.CompoundID, compound.CompoundAccession, compound.CompoundName, compound.CompoundHTMLName, compound.Accession)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		// Add compound.Access to the compoundKeys map
		compoundKeys[compound.Accession] = true
	}

	// Next, insert the ReactionSides and ReactionParticipants
	for _, reactionParticipant := range rhea.ReactionParticipants {
		// Insert ReactionSide, which is needed to insert the ReactionParticipant
		_, err = tx.Exec("INSERT INTO reactionside(accession) VALUES (?) ON CONFLICT DO NOTHING", reactionParticipant.ReactionSide)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// Insert the ReactionParticipants
		_, err = tx.Exec("INSERT INTO reactionparticipant(reactionside, contains, containsn, minus, plus, compound) VALUES (?, ?, ?, ?, ?, ?)", reactionParticipant.ReactionSide, reactionParticipant.Contains, reactionParticipant.ContainsN, reactionParticipant.Minus, reactionParticipant.Plus, reactionParticipant.Compound)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	// Insert the Reactions themselves
	for _, reaction := range rhea.Reactions {
		// Insert Reaction
		_, err = tx.Exec("INSERT INTO reaction(id, directional, accession, status, comment, equation, htmlequation, ischemicallybalanced, istransport, ec, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", reaction.ID, reaction.Directional, reaction.Accession, reaction.Status, reaction.Comment, reaction.Equation, reaction.HTMLEquation, reaction.IsChemicallyBalanced, reaction.IsTransport, reaction.Ec, reaction.Location)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		// Insert ReactionsideReaction. Basically, these represent the substrates, products, or substratesOrProducts of a given reaction
		for _, substrate := range reaction.Substrates {
			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'substrate')", reaction.Accession, substrate)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		for _, product := range reaction.Products {
			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'product')", reaction.Accession, product)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		for _, substrateorproduct := range reaction.SubstrateOrProducts {
			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'substrateorproduct')", reaction.Accession, substrateorproduct)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func RheaTsvInsert(db *sqlx.DB, path string, gzipped bool) error {
	lines := make(chan rhea.RheaToUniprot, 10000)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	// Handle gzipped TSV insert
	if gzipped {
		r, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		go rhea.ParseRheaToUniprotTsv(r, lines)
	} else {
		go rhea.ParseRheaToUniprotTsv(file, lines)
	}
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	var counter int
	for {
		line, more := <-lines
		if more {
			counter++
			var accession string
			err = tx.Get(&accession, "SELECT accession FROM reaction WHERE id = ?", line.RheaID)
			if err == nil {
				_, err := tx.Exec("INSERT OR IGNORE INTO uniprot_to_reaction(reaction,uniprot) VALUES (?,?)", accession, line.UniprotID)
				if err != nil {
					_ = tx.Rollback()
					return err
				}
				if counter >= 10000 {
					counter = 0
					err = tx.Commit()
					if err != nil {
						_ = tx.Rollback()
						return err
					}
					tx, err = db.Beginx()
					if err != nil {
						return err
					}
				}
			} else {
				fmt.Println(line.RheaID)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			return nil
		}
	}
}

/******************************************************************************

Uniprot

we should set this to operate like the above rheatouniprot tsv stuff

******************************************************************************/

func UniprotInsert(db *sqlx.DB, uniprotDatabase string, entries chan uniprot.Entry, errors chan error, wg *sync.WaitGroup) {
	var counter int
	var err error
	defer wg.Done()
	tx, err := db.Begin()
	if err != nil {
		errors <- err
	}
	for {
		entry, more := <-entries
		if more {
			counter++
			// Insert seqhashes
			sequence := strings.ToUpper(entry.Sequence.Value)
			seqhash, err := seqhash.Hash(sequence, "PROTEIN", false, false)
			if err != nil {
				errors <- err
				continue
			}
			_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", seqhash, sequence, false, false, "PROTEIN")
			if err != nil {
				errors <- err
				continue
			}

			// Insert uniprot
			_, err = tx.Exec("INSERT INTO uniprot(database,accession,seqhash) VALUES (?,?,?)", uniprotDatabase, entry.Accession[0], seqhash)
			if err != nil {
				errors <- err
				continue
			}
			if counter >= 10000 {
				counter = 0
				err = tx.Commit()
				if err != nil {
					errors <- err
					continue
				}
				tx, err = db.Begin()
				if err != nil {
					errors <- err
					continue
				}
			}
		} else {
			err = tx.Commit()
			if err != nil {
				errors <- err
			}
			return
		}
	}
}

/******************************************************************************

Genbank

******************************************************************************/

func GenbankInsert(db *sqlx.DB, genbankList []poly.Sequence) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, gbk := range genbankList {
		// FIX THIS
		sq, err := seqhash.Hash(gbk.Sequence, "DNA", false, true)
		if err != nil {
			return err
		}
		var sequenceType string
		var circular bool
		var doubleStranded bool
		switch sq[3] {
		case 'D':
			sequenceType = "DNA"
		case 'R':
			sequenceType = "RNA"
		}

		switch sq[4] {
		case 'L':
			circular = false
		case 'C':
			circular = true
		}

		switch sq[5] {
		case 'L':
			doubleStranded = false
		case 'D':
			doubleStranded = true
		}

		// Insert initial seqhash
		_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", sq, strings.ToUpper(gbk.Sequence), circular, doubleStranded, sequenceType)
		if err != nil {
			return err
		}

		// Insert genbank file
		_, err = tx.Exec("INSERT INTO genbank(accession, seqhash) VALUES (?,?)", gbk.Meta.Locus.Name, sq)
		if err != nil {
			return err
		}

		// For each protein, insert seqhash of the protein, insert hash of the current gene, and then add in a genbankfeatures
		for _, feature := range gbk.Features {
			if feature.Type == "CDS" {
				translation := feature.Attributes["translation"]
				if translation != "" {
					// Insert protein seqhash
					proteinSeqhash, err := seqhash.Hash(translation, "PROTEIN", false, false)
					if err != nil {
						return err
					}
					_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", proteinSeqhash, translation, false, false, "PROTEIN")
					if err != nil {
						return err
					}
					// Insert gene seqhash
					geneSequence := feature.GetSequence()
					geneSeqhash, err := seqhash.Hash(geneSequence, sequenceType, false, true)
					if err != nil {
						return err
					}
					_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype,translation) VALUES(?,?,?,?,?,?)", geneSeqhash, geneSequence, false, true, sequenceType, proteinSeqhash)
					if err != nil {
						return err
					}
					// Insert reference to the feature
					_, err = tx.Exec("INSERT INTO genbankfeatures(seqhash,genbank) VALUES (?,?)", geneSeqhash, gbk.Meta.Locus.Name)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

/******************************************************************************

Chembl

******************************************************************************/

func ChemblAttach(db *sqlx.DB, chembl string) error {
	_, err := db.Exec("ATTACH DATABASE ? AS chembl", chembl)
	if err != nil {
		return err
	}
	return nil
}
