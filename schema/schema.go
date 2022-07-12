package schema

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

// CreateDatabase creates a new database with the given name.
func CreateDatabase(config config.Config) error {

	// Begin SQLite
	log.Println("Creating database...")
	db, err := sqlx.Open("sqlite", config.AllbasePath)

	if err != nil {
		log.Fatalf("Failed to open sqlite database: %s", err)
	}

	defer db.Close()

	// Execute our schema in memory
	_, err = db.Exec(CreateSchema())
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	schemaStringBytes, err := ioutil.ReadFile(config.ChemblSchema)
	if err != nil {
		log.Fatalf("Failed to open chembl schema: %s", err)
	}

	_, err = db.Exec(string(schemaStringBytes))
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	err = chemblAttach(db, config.AllbasePath)
	if err != nil {
		log.Fatalf("Failed to attach chembl with error %s", err)
	}

	return err
}

// CreateSchema generates the SQL for the database schema minus attachments which are handled in createDatabase.
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
		HTMLNAME                       = "html_name"
		REACTION                       = "reaction"
		REACTIONSIDE                   = "reactionside"
		UNIPROT                        = "uniprot"
		ALLBASEDOT                     = ""
		ALLBASE                        = "allbase"
	)

	// each built string will be appended to this slice and returned at the end of the function
	var tableStringSlice []string

	// // create the allbase database itself
	// databaseDeclaration := "CREATE DATABASE " + ALLBASE
	// tableStringSlice = append(tableStringSlice, databaseDeclaration)

	// create seqhash table
	seqhash := sqlbuilder.NewCreateTableBuilder()
	seqhash.CreateTable(SEQHASH).IfNotExists()
	seqhash.Define(SEQHASH, TEXT, NOTNULL, PRIMARYKEY)
	seqhash.Define("sequence", TEXT, NOTNULL)
	seqhash.Define("circular", BOOL, NOTNULL, DEFAULTFALSE)
	seqhash.Define("doublestranded", BOOL, NOTNULL, DEFAULTTRUE)
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
	genbankFeatures := sqlbuilder.NewCreateTableBuilder()
	genbankFeatures.CreateTable("genbank_features").IfNotExists()
	genbankFeatures.Define(SEQHASH, TEXT, NOTNULL, REFERENCESEQHASH)
	genbankFeatures.Define("parent", TEXT, NOTNULL, "REFERENCES genbank(accession)")
	genbankFeatures.Define("PRIMARY KEY(", SEQHASH, ", ", "parent", ")")
	genbankFeaturesTableString, _ := genbankFeatures.Build()
	tableStringSlice = append(tableStringSlice, genbankFeaturesTableString)

	// create uniprot table
	uniprot := sqlbuilder.NewCreateTableBuilder()
	uniprot.CreateTable(UNIPROT).IfNotExists()
	uniprot.Define(ACCESSION, TEXT, PRIMARYKEY)
	uniprot.Define("database", TEXT, NOTNULL)
	uniprot.Define(SEQHASH, TEXT, NOTNULL, REFERENCESEQHASH)
	uniprotTableString, _ := uniprot.Build()
	tableStringSlice = append(tableStringSlice, uniprotTableString)

	//*** create rhea tables ***//

	// create chebi table
	chebi := sqlbuilder.NewCreateTableBuilder()
	chebi.CreateTable(CHEBI).IfNotExists()
	chebi.Define(ACCESSION, TEXT, PRIMARYKEY)
	chebi.Define("subclass_of", TEXT, REFERENCECHEBIACCESSION)
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
	compound.Define("polymerization_index", TEXT)
	compound.Define("compound_type", TEXT, NOTNULL, "CHECK(compound_type IN ('SmallMolecule', 'Polymer', 'GenericPolypeptide', 'GenericPolynucleotide', 'GenericHeteropolysaccharide'))")
	compoundTableString, _ := compound.Build()
	tableStringSlice = append(tableStringSlice, compoundTableString)

	// create reactivePart table
	reactivePart := sqlbuilder.NewCreateTableBuilder()
	reactivePart.CreateTable("reactive_part").IfNotExists()
	reactivePart.Define(ID, INTEGER)
	reactivePart.Define(ACCESSION, TEXT, PRIMARYKEY)
	reactivePart.Define(NAME, TEXT)
	reactivePart.Define(HTMLNAME, TEXT)
	reactivePart.Define(COMPOUND, TEXT, NOTNULL, REFERENCECOMPOUNDACCESSION)
	reactivePartTableString, _ := reactivePart.Build()
	tableStringSlice = append(tableStringSlice, reactivePartTableString)

	// create reaction table
	reaction := sqlbuilder.NewCreateTableBuilder()
	reaction.CreateTable(REACTION).IfNotExists()
	reaction.Define(ID, INTEGER)
	reaction.Define("directional", BOOL, NOTNULL, DEFAULTFALSE)
	reaction.Define(ACCESSION, TEXT, PRIMARYKEY)
	reaction.Define("status", TEXT)
	reaction.Define("comment", TEXT)
	reaction.Define("equation", TEXT)
	reaction.Define("html_equation", TEXT)
	reaction.Define("is_chemically_balanced", BOOL, NOTNULL, DEFAULTTRUE)
	reaction.Define("is_transport", BOOL, NOTNULL, DEFAULTFALSE)
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

	// create reactionside_reaction table
	reactionsideReaction := sqlbuilder.NewCreateTableBuilder()
	reactionsideReaction.CreateTable("reactionside_reaction").IfNotExists()
	reactionsideReaction.Define(REACTION, TEXT, NOTNULL, REFERENCEREACTIONACCESSION)
	reactionsideReaction.Define(REACTIONSIDE, TEXT, NOTNULL, REFERENCEREACTIONSIDEACCESSION)
	reactionsideReaction.Define("reactionside_reaction_type", TEXT, NOTNULL, "CHECK(reactionside_reaction_type IN ('substrate_or_product', 'substrate', 'product'))")
	reactionsideReaction.Define("PRIMARY KEY(", REACTION, ", ", REACTIONSIDE, ")")
	reactionsideReactionTableString, _ := reactionsideReaction.Build()
	tableStringSlice = append(tableStringSlice, reactionsideReactionTableString)

	// create reactionParticipant table
	reactionParticipant := sqlbuilder.NewCreateTableBuilder()
	reactionParticipant.CreateTable("reaction_participant").IfNotExists()
	reactionParticipant.Define(COMPOUND, TEXT, REFERENCECOMPOUNDACCESSION)
	reactionParticipant.Define(REACTIONSIDE, TEXT, NOTNULL, REFERENCEREACTIONSIDEACCESSION)
	reactionParticipant.Define("contains", INTEGER)
	reactionParticipant.Define("contains_n", BOOL, NOTNULL, DEFAULTFALSE)
	reactionParticipant.Define("minus", BOOL, NOTNULL, DEFAULTFALSE)
	reactionParticipant.Define("plus", BOOL, NOTNULL, DEFAULTFALSE)
	reactionParticipant.Define("PRIMARY KEY(", COMPOUND, ", ", REACTIONSIDE, ")")
	reactionparticipantTableString, _ := reactionParticipant.Build()
	tableStringSlice = append(tableStringSlice, reactionparticipantTableString)

	// create uniprot_to_reaction table
	uniprotToReaction := sqlbuilder.NewCreateTableBuilder()
	uniprotToReaction.CreateTable("uniprot_to_reaction").IfNotExists()
	uniprotToReaction.Define(REACTION, TEXT, REFERENCEREACTIONACCESSION)
	uniprotToReaction.Define(UNIPROT, TEXT, "REFERENCES uniprot(accession)")
	uniprotToReaction.Define("PRIMARY KEY(", REACTION, ", ", UNIPROT, ")")
	uniprotToReactionTableString, _ := uniprotToReaction.Build()
	tableStringSlice = append(tableStringSlice, uniprotToReactionTableString)

	schema := strings.Join(tableStringSlice, ";\n\n")

	return schema
}

/******************************************************************************

Chembl

******************************************************************************/

func chemblAttach(db *sqlx.DB, chembl string) error {
	_, err := db.Exec("ATTACH DATABASE ? AS chembl", chembl)
	if err != nil {
		return err
	}
	return nil
}
