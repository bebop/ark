package schema

import (
	"log"
	"os"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func createDatabase(dbPath string) error {
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		log.Fatal("Database already exists. Run 'allbase clean' to remove it.")
	}

	// Begin SQLite
	log.Println("Creating database...")
	db, err := sqlx.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatalf("Failed to open sqlite in %s: %s", dbPath, err)
	}

	defer db.Close()

	// Execute our schema in memory
	_, err = db.Exec(CreateSchema())
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	return nil
}

// CreateSchema generates the SQL for the database schema
func CreateSchema() string {

	// Note:
	// Some variables have wonky capitalizations.
	// This is because the SQLite driver is case-sensitive
	// and we want to be consistent with the rest of the function.

	// Frequenty used strings in schema definition defined here
	// for convenience and to avoid typos.
	const (
		TEXT                           = "TEXT"
		INTEGER                        = "INT"
		BOOL                           = "BOOL"
		NOTNULL                        = "NOT NULL"
		PRIMARYKEY                     = "PRIMARY KEY"
		DEFAULTFALSE                   = "DEFAULT FALSE"
		DEFAULTTRUE                    = "DEFAULT TRUE"
		SEQHASH                        = "seqhash"
		ACCESSION                      = "accession"
		GENBANK                        = "genbank"
		REFERENCESEQHASH               = "REFERENCES seqhash(seqhash)"
		REFERENCECHEBIACCESSION        = "REFERENCES chebi(accession)"
		REFERENCECOMPOUNDACCESSION     = "REFERENCES compound(accession)"
		REFERENCEREACTIONACCESSION     = "REFERENCES reaction(accession)"
		REFERENCEREACTIONSIDEACCESSION = "REFERENCES reactionside(accession)"
		ID                             = "id"
		NAME                           = "name"
		COMPOUND                       = "compound"
		CHEBI                          = "chebi"
		HTMLNAME                       = "htmlname"
		REACTION                       = "reaction"
		REACTIONSIDE                   = "reactionside"
		UNIPROT                        = "uniprot"
	)

	// each built string will be appended to this slice and returned at the end of the function
	var tableStringSlice []string

	// create seqhash table
	seqhash := sqlbuilder.NewCreateTableBuilder()
	seqhash.CreateTable(SEQHASH).IfNotExists()
	seqhash.Define(SEQHASH, TEXT, NOTNULL, PRIMARYKEY)
	seqhash.Define("sequence", TEXT, NOTNULL)
	seqhash.Define("circular", INTEGER, NOTNULL, DEFAULTFALSE)
	seqhash.Define("doublestranded", INTEGER, NOTNULL, DEFAULTTRUE)
	seqhash.Define("seqhashtype", TEXT, NOTNULL, "CHECK (seqhashtype IN ('DNA', 'RNA', 'PROTEIN'))")
	seqhash.Define("translation", TEXT, REFERENCESEQHASH)
	seqhashTableString, _ := seqhash.Build()
	tableStringSlice = append(tableStringSlice, seqhashTableString)
	// "CREATE TABLE IF NOT EXISTS seqhash (seqhash TEXT NOT NULL PRIMARY KEY, sequence TEXT NOT NULL, circular INTEGER NOT NULL DEFAULT FALSE, doublestranded INTEGER NOT NULL DEFAULT TRUE, seqhashtype TEXT NOT NULL CHECK (seqhashtype IN ('DNA', 'RNA', 'PROTEIN')), translations TEXT REFERENCES seqhash(seqhash))"

	// create genbank table
	genbank := sqlbuilder.NewCreateTableBuilder()
	genbank.CreateTable(GENBANK).IfNotExists()
	genbank.Define(ACCESSION, TEXT, PRIMARYKEY)
	genbank.Define(SEQHASH, TEXT, NOTNULL, REFERENCESEQHASH)
	genbankTableString, _ := genbank.Build()
	tableStringSlice = append(tableStringSlice, genbankTableString)

	// create genbank features table
	genbankfeatures := sqlbuilder.NewCreateTableBuilder()
	genbankfeatures.CreateTable("genbankfeatures").IfNotExists()
	genbankfeatures.Define(SEQHASH, TEXT, NOTNULL, REFERENCESEQHASH)
	genbankfeatures.Define(GENBANK, TEXT, NOTNULL, "REFERENCES genbank(accession)")
	genbankfeatures.Define("PRIMARY KEY(", SEQHASH, ", ", GENBANK, ")")
	genbankfeaturesTableString, _ := genbankfeatures.Build()
	tableStringSlice = append(tableStringSlice, genbankfeaturesTableString)

	// create uniprot table
	uniprot := sqlbuilder.NewCreateTableBuilder()
	uniprot.CreateTable(UNIPROT).IfNotExists()
	uniprot.Define(ACCESSION, TEXT, PRIMARYKEY)
	uniprot.Define("database", TEXT, NOTNULL)
	uniprot.Define(SEQHASH, TEXT, NOTNULL, REFERENCESEQHASH)
	uniprotTableString, _ := uniprot.Build()
	tableStringSlice = append(tableStringSlice, uniprotTableString)

	//*** create rhea tables ***//

	// create chebi table <- what is chebi @Koeng101? chembl?
	chebi := sqlbuilder.NewCreateTableBuilder()
	chebi.CreateTable(CHEBI).IfNotExists()
	chebi.Define(ACCESSION, TEXT, PRIMARYKEY)
	chebi.Define("subclassof", TEXT, REFERENCECHEBIACCESSION)
	chebiTableString, _ := chebi.Build()
	tableStringSlice = append(tableStringSlice, chebiTableString)

	// create compound table
	compound := sqlbuilder.NewCreateTableBuilder()
	compound.CreateTable(COMPOUND).IfNotExists()
	compound.Define(ID, INTEGER, NOTNULL)
	compound.Define(ACCESSION, TEXT, PRIMARYKEY)
	compound.Define("position", TEXT)
	compound.Define(NAME, TEXT)
	compound.Define(HTMLNAME, TEXT)
	compound.Define("formula", TEXT)
	compound.Define("charge", TEXT)
	compound.Define(CHEBI, TEXT, REFERENCECHEBIACCESSION)
	compound.Define("polymerizationindex", TEXT)
	compound.Define("compoundtype", TEXT, NOTNULL, "CHECK(compoundtype IN ('SmallMolecule', 'Polymer', 'GenericPolypeptide', 'GenericPolynucleotide', 'GenericHeteropolysaccharide'))")
	compoundTableString, _ := compound.Build()
	tableStringSlice = append(tableStringSlice, compoundTableString)

	// create reactivepart table
	reactivepart := sqlbuilder.NewCreateTableBuilder()
	reactivepart.CreateTable("reactivepart").IfNotExists()
	reactivepart.Define(ID, INTEGER)
	reactivepart.Define(ACCESSION, TEXT, PRIMARYKEY)
	reactivepart.Define(NAME, TEXT)
	reactivepart.Define(HTMLNAME, TEXT)
	reactivepart.Define(COMPOUND, TEXT, NOTNULL, REFERENCECOMPOUNDACCESSION)
	reactivepartTableString, _ := reactivepart.Build()
	tableStringSlice = append(tableStringSlice, reactivepartTableString)

	// create reaction table
	reaction := sqlbuilder.NewCreateTableBuilder()
	reaction.CreateTable(REACTION).IfNotExists()
	reaction.Define(ID, INTEGER)
	reaction.Define("directional", BOOL, NOTNULL, DEFAULTFALSE)
	reaction.Define(ACCESSION, TEXT, PRIMARYKEY)
	reaction.Define("status", TEXT)
	reaction.Define("comment", TEXT)
	reaction.Define("equation", TEXT)
	reaction.Define("htmlequation", TEXT)
	reaction.Define("ischemicallybalanced", BOOL, NOTNULL, DEFAULTTRUE)
	reaction.Define("istransport", BOOL, NOTNULL, DEFAULTFALSE)
	reaction.Define("ec", TEXT)
	reaction.Define("location", TEXT)
	reactionTableString, _ := reaction.Build()
	tableStringSlice = append(tableStringSlice, reactionTableString)

	// create reactionside table
	reactionside := sqlbuilder.NewCreateTableBuilder()
	reactionside.CreateTable(REACTIONSIDE).IfNotExists()
	reactionside.Define(ACCESSION, TEXT, PRIMARYKEY)
	reactionsideTableString, _ := reactionside.Build()
	tableStringSlice = append(tableStringSlice, reactionsideTableString)

	// create reactionsidereaction table
	reactionsidereaction := sqlbuilder.NewCreateTableBuilder()
	reactionsidereaction.CreateTable("reactionsidereaction").IfNotExists()
	reactionsidereaction.Define(REACTION, TEXT, NOTNULL, REFERENCEREACTIONACCESSION)
	reactionsidereaction.Define(REACTIONSIDE, TEXT, NOTNULL, REFERENCEREACTIONSIDEACCESSION)
	reactionsidereaction.Define("reactionsidereactiontype", TEXT, NOTNULL, "CHECK(reactionsidereactiontype IN ('substrateorproduct', 'substrate', 'product'))")
	reactionsidereaction.Define("PRIMARY KEY(", REACTION, ", ", REACTIONSIDE, ")")
	reactionsidereactionTableString, _ := reactionsidereaction.Build()
	tableStringSlice = append(tableStringSlice, reactionsidereactionTableString)

	// create reactionparticipant table
	reactionparticipant := sqlbuilder.NewCreateTableBuilder()
	reactionparticipant.CreateTable("reactionparticipant").IfNotExists()
	reactionparticipant.Define(COMPOUND, TEXT, REFERENCECOMPOUNDACCESSION)
	reactionparticipant.Define(REACTIONSIDE, TEXT, NOTNULL, REFERENCEREACTIONSIDEACCESSION)
	reactionparticipant.Define("contains", INTEGER)
	reactionparticipant.Define("containsn", BOOL, NOTNULL, DEFAULTFALSE)
	reactionparticipant.Define("minus", BOOL, NOTNULL, DEFAULTFALSE)
	reactionparticipant.Define("plus", BOOL, NOTNULL, DEFAULTFALSE)
	reactionparticipant.Define("PRIMARY KEY(", COMPOUND, ", ", REACTIONSIDE, ")")
	reactionparticipantTableString, _ := reactionparticipant.Build()
	tableStringSlice = append(tableStringSlice, reactionparticipantTableString)

	// create uniprot_to_reaction table
	uniprotToReaction := sqlbuilder.NewCreateTableBuilder()
	uniprotToReaction.CreateTable("uniprot_to_reaction").IfNotExists()
	uniprotToReaction.Define(REACTION, TEXT, REFERENCEREACTIONACCESSION)
	uniprotToReaction.Define(UNIPROT, TEXT, "REFERENCES uniprot(accession)")
	uniprotToReaction.Define("PRIMARY KEY(", REACTION, ", ", UNIPROT, ")")
	uniprotToReactionTableString, _ := uniprotToReaction.Build()
	tableStringSlice = append(tableStringSlice, uniprotToReactionTableString)

	// return schema as string slice where each element is a table string
	schema := strings.Join(tableStringSlice, ";\n\n")

	return schema
}

/******************************************************************************

Schema structs begin here

******************************************************************************/

// type Seqhash struct {
// 	Seqhash        string         `db:"seqhash"`
// 	Circular       bool           `db:"circular"`
// 	DoubleStranded bool           `db:"doublestranded"`
// 	Sequence       string         `db:"sequence"`
// 	SeqhashType    string         `db:"seqhashtype"`
// 	Translation    sql.NullString `db:"translation"`
// }

// type Genbank struct {
// 	Id          string `db:"id"`
// 	Genbank     string `db:"genbank"`
// 	GenbankHash string `db:"genbankhash"`
// 	PolyVersion string `db:"polyversion"`
// 	Seqhash     string `db:"seqhash"`
// }

// type GenbankFeature struct {
// 	Seqhash string `db:"seqhash"`
// 	Genbank string `db:"genbank"`
// }

// type Uniprot struct {
// 	Accession   string         `db:"accession"`
// 	UniprotHash string         `db:"uniprothash"`
// 	Uniprot     types.JSONText `db:"uniprot"`
// 	Seqhash     string         `db:"seqhash"`
// }

// type GenbankFull struct {
// 	Seqhashes       []Seqhash
// 	Genbank         Genbank
// 	GenbankFeatures []GenbankFeature
// }

/******************************************************************************

Rhea

******************************************************************************/

// RheaInsert inserts the Rhea database into an SQLite database with proper normalization for advanced queries.
// func RheaInsert(db *sqlx.DB, rhea rhea.Rhea) error {
// 	// Start transaction with database for insertion. This ensures if there are any problems, they are seamlessly rolled back
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	// First, insert ChEBIs and Compounds
// 	compoundKeys := make(map[string]bool)
// 	for _, compound := range rhea.Compounds {
// 		// Insert root chebi. Ie, what this current compound's subclass is
// 		_, err = tx.Exec("INSERT OR IGNORE INTO chebi(accession) VALUES (?)", compound.SubclassOfChEBI)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}
// 		// Insert the chebi of the current compound. If it is already inserted, update the subclassification
// 		_, err = tx.Exec("INSERT INTO chebi(accession, subclassof) VALUES (?, ?) ON CONFLICT (accession) DO UPDATE SET subclassof = ?", compound.ChEBI, compound.SubclassOfChEBI, compound.SubclassOfChEBI)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}
// 		// Insert the compound itself
// 		_, err = tx.Exec("INSERT INTO compound(id, accession, position, name, htmlname, formula, charge, chebi, polymerizationindex, compoundtype) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING", compound.ID, compound.Accession, compound.Position, compound.Name, compound.HTMLName, compound.Formula, compound.Charge, compound.ChEBI, compound.PolymerizationIndex, compound.CompoundType)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}
// 		// If the compound isn't a small molecule or polymer, that means it would be a reactive part of a larger compound. So we add it in
// 		if (compound.CompoundType != "SmallMolecule") && (compound.CompoundType != "Polymer") {
// 			_, err = tx.Exec("INSERT INTO reactivepart(id, accession, name, htmlname, compound) VALUES (?, ?, ?, ?, ?)", compound.CompoundID, compound.CompoundAccession, compound.CompoundName, compound.CompoundHTMLName, compound.Accession)
// 			if err != nil {
// 				_ = tx.Rollback()
// 				return err
// 			}
// 		}
// 		// Add compound.Access to the compoundKeys map
// 		compoundKeys[compound.Accession] = true
// 	}

// 	// Next, insert the ReactionSides and ReactionParticipants
// 	for _, reactionParticipant := range rhea.ReactionParticipants {
// 		// Insert ReactionSide, which is needed to insert the ReactionParticipant
// 		_, err = tx.Exec("INSERT INTO reactionside(accession) VALUES (?) ON CONFLICT DO NOTHING", reactionParticipant.ReactionSide)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}
// 		// Insert the ReactionParticipants
// 		_, err = tx.Exec("INSERT INTO reactionparticipant(reactionside, contains, containsn, minus, plus, compound) VALUES (?, ?, ?, ?, ?, ?)", reactionParticipant.ReactionSide, reactionParticipant.Contains, reactionParticipant.ContainsN, reactionParticipant.Minus, reactionParticipant.Plus, reactionParticipant.Compound)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}
// 	}

// 	// Insert the Reactions themselves
// 	for _, reaction := range rhea.Reactions {
// 		// Insert Reaction
// 		_, err = tx.Exec("INSERT INTO reaction(id, directional, accession, status, comment, equation, htmlequation, ischemicallybalanced, istransport, ec, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", reaction.ID, reaction.Directional, reaction.Accession, reaction.Status, reaction.Comment, reaction.Equation, reaction.HTMLEquation, reaction.IsChemicallyBalanced, reaction.IsTransport, reaction.Ec, reaction.Location)
// 		if err != nil {
// 			_ = tx.Rollback()
// 			return err
// 		}

// 		// Insert ReactionsideReaction. Basically, these represent the substrates, products, or substratesOrProducts of a given reaction
// 		for _, substrate := range reaction.Substrates {
// 			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'substrate')", reaction.Accession, substrate)
// 			if err != nil {
// 				_ = tx.Rollback()
// 				return err
// 			}
// 		}
// 		for _, product := range reaction.Products {
// 			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'product')", reaction.Accession, product)
// 			if err != nil {
// 				_ = tx.Rollback()
// 				return err
// 			}
// 		}
// 		for _, substrateorproduct := range reaction.SubstrateOrProducts {
// 			_, err = tx.Exec("INSERT INTO reactionsidereaction(reaction, reactionside, reactionsidereactiontype) VALUES (?, ?, 'substrateorproduct')", reaction.Accession, substrateorproduct)
// 			if err != nil {
// 				_ = tx.Rollback()
// 				return err
// 			}
// 		}
// 	}

// 	// Commit
// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func RheaTsvInsert(db *sqlx.DB, path string, gzipped bool) error {
// 	lines := make(chan rhea.RheaToUniprot, 10000)
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	// Handle gzipped TSV insert
// 	if gzipped {
// 		r, err := gzip.NewReader(file)
// 		if err != nil {
// 			return err
// 		}
// 		go rhea.ParseRheaToUniprotTsv(r, lines)
// 	} else {
// 		go rhea.ParseRheaToUniprotTsv(file, lines)
// 	}
// 	tx, err := db.Beginx()
// 	if err != nil {
// 		return err
// 	}
// 	var counter int
// 	for {
// 		line, more := <-lines
// 		if more {
// 			counter++
// 			var accession string
// 			err = tx.Get(&accession, "SELECT accession FROM reaction WHERE id = ?", line.RheaID)
// 			if err == nil {
// 				_, err := tx.Exec("INSERT OR IGNORE INTO uniprot_to_reaction(reaction,uniprot) VALUES (?,?)", accession, line.UniprotID)
// 				if err != nil {
// 					_ = tx.Rollback()
// 					return err
// 				}
// 				if counter >= 10000 {
// 					counter = 0
// 					err = tx.Commit()
// 					if err != nil {
// 						_ = tx.Rollback()
// 						return err
// 					}
// 					tx, err = db.Beginx()
// 					if err != nil {
// 						return err
// 					}
// 				}
// 			} else {
// 				fmt.Println(line.RheaID)
// 			}
// 		} else {
// 			err = tx.Commit()
// 			if err != nil {
// 				_ = tx.Rollback()
// 				return err
// 			}
// 			return nil
// 		}
// 	}
// }

// /******************************************************************************

// Uniprot

// we should set this to operate like the above rheatouniprot tsv stuff

// ******************************************************************************/

// func UniprotInsert(db *sqlx.DB, uniprotDatabase string, entries chan uniprot.Entry, errors chan error, wg *sync.WaitGroup) {
// 	var counter int
// 	var err error
// 	defer wg.Done()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		errors <- err
// 	}
// 	for {
// 		entry, more := <-entries
// 		if more {
// 			counter++
// 			// Insert seqhashes
// 			sequence := strings.ToUpper(entry.Sequence.Value)
// 			seqhash, err := seqhash.Hash(sequence, "PROTEIN", false, false)
// 			if err != nil {
// 				errors <- err
// 				continue
// 			}
// 			_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", seqhash, sequence, false, false, "PROTEIN")
// 			if err != nil {
// 				errors <- err
// 				continue
// 			}

// 			// Insert uniprot
// 			_, err = tx.Exec("INSERT INTO uniprot(database,accession,seqhash) VALUES (?,?,?)", uniprotDatabase, entry.Accession[0], seqhash)
// 			if err != nil {
// 				errors <- err
// 				continue
// 			}
// 			if counter >= 10000 {
// 				counter = 0
// 				err = tx.Commit()
// 				if err != nil {
// 					errors <- err
// 					continue
// 				}
// 				tx, err = db.Begin()
// 				if err != nil {
// 					errors <- err
// 					continue
// 				}
// 			}
// 		} else {
// 			err = tx.Commit()
// 			if err != nil {
// 				errors <- err
// 			}
// 			return
// 		}
// 	}
// }

// /******************************************************************************

// Genbank

// ******************************************************************************/

// func GenbankInsert(db *sqlx.DB, genbankList []poly.Sequence) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	for _, gbk := range genbankList {
// 		// FIX THIS
// 		sq, err := seqhash.Hash(gbk.Sequence, "DNA", false, true)
// 		if err != nil {
// 			return err
// 		}
// 		var sequenceType string
// 		var circular bool
// 		var doubleStranded bool
// 		switch sq[3] {
// 		case 'D':
// 			sequenceType = "DNA"
// 		case 'R':
// 			sequenceType = "RNA"
// 		}

// 		switch sq[4] {
// 		case 'L':
// 			circular = false
// 		case 'C':
// 			circular = true
// 		}

// 		switch sq[5] {
// 		case 'L':
// 			doubleStranded = false
// 		case 'D':
// 			doubleStranded = true
// 		}

// 		// Insert initial seqhash
// 		_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", sq, strings.ToUpper(gbk.Sequence), circular, doubleStranded, sequenceType)
// 		if err != nil {
// 			return err
// 		}

// 		// Insert genbank file
// 		_, err = tx.Exec("INSERT INTO genbank(accession, seqhash) VALUES (?,?)", gbk.Meta.Locus.Name, sq)
// 		if err != nil {
// 			return err
// 		}

// 		// For each protein, insert seqhash of the protein, insert hash of the current gene, and then add in a genbankfeatures
// 		for _, feature := range gbk.Features {
// 			if feature.Type == "CDS" {
// 				translation := feature.Attributes["translation"]
// 				if translation != "" {
// 					// Insert protein seqhash
// 					proteinSeqhash, err := seqhash.Hash(translation, "PROTEIN", false, false)
// 					if err != nil {
// 						return err
// 					}
// 					_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype) VALUES (?,?,?,?,?)", proteinSeqhash, translation, false, false, "PROTEIN")
// 					if err != nil {
// 						return err
// 					}
// 					// Insert gene seqhash
// 					geneSequence := feature.GetSequence()
// 					geneSeqhash, err := seqhash.Hash(geneSequence, sequenceType, false, true)
// 					if err != nil {
// 						return err
// 					}
// 					_, err = tx.Exec("INSERT OR IGNORE INTO seqhash(seqhash,sequence,circular,doublestranded,seqhashtype,translation) VALUES(?,?,?,?,?,?)", geneSeqhash, geneSequence, false, true, sequenceType, proteinSeqhash)
// 					if err != nil {
// 						return err
// 					}
// 					// Insert reference to the feature
// 					_, err = tx.Exec("INSERT INTO genbankfeatures(seqhash,genbank) VALUES (?,?)", geneSeqhash, gbk.Meta.Locus.Name)
// 					if err != nil {
// 						return err
// 					}
// 				}
// 			}
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// /******************************************************************************

// Chembl

// ******************************************************************************/

// func ChemblAttach(db *sqlx.DB, chembl string) error {
// 	_, err := db.Exec("ATTACH DATABASE ? AS chembl", chembl)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
