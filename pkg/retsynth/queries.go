package retsynth

import (
	"math"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Easy database connector
func ConnectDB() (*sqlx.DB, error) {

	// Get the file path of the Retsynth database from the environment variable, if it exists otherwise set default path
	var RetsynthDBPath, ok = os.LookupEnv("RETSYNTH_DB_PATH")
	if !ok {
		RetsynthDBPath = "../../data/dev/retsynth/minimal.db"
	}
	var db *sqlx.DB
	var err error
	db, err = sqlx.Connect("sqlite3", RetsynthDBPath)
	if err != nil {
		return db, err
	}
	return db, err
}

// Retrieves unique metabolic clusters (organisms with the exact same metabolism) in the database
func GetUniqueMetabolicClusters() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var clusters []string
	query := "SELECT DISTINCT cluster_num FROM cluster"
	err = db.Select(&clusters, query)
	if err != nil {
		return nil, err
	}
	return clusters, err
}

// Retrieves model IDs from a specified cluster in the database
func GetModelIDsFromCluster(cluster string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var models []string
	query := "SELECT model_id FROM cluster WHERE cluster_num = ?"
	err = db.Select(&models, query, cluster)
	if err != nil {
		return nil, err
	}
	return models, err
}

// Retrieves all model IDs from the database
func GetAllModelIDs() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var models []string
	query := "SELECT DISTINCT model_id FROM model"
	err = db.Select(&models, query)
	if err != nil {
		return nil, err
	}
	return models, err
}

// Retrieves all model objects from the database
func GetAllModels() []Model {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	var models []Model
	query := "SELECT * FROM model"
	err = db.Select(&models, query)
	if err != nil {
		panic(err)
	}
	return models
}

// Retrieves name of organism given a specific organism ID
func GetOrganismName(organismID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var name string
	query := "SELECT name FROM organism WHERE ID = ?"
	err = db.Get(&name, query, organismID)
	if err != nil {
		return "", err
	}
	return name, err
}

// Retrieves ID of organism given a specific organism name
func GetOrganismID(organismName string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var ID string
	query := "SELECT ID FROM organism WHERE name = ?"
	err = db.Get(&ID, query, organismName)
	if err != nil {
		return "", err
	}
	return ID, err
}

// Retrieves compound ID given a compound name
func GetCompoundID(compoundName string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var ID string
	query := "SELECT ID FROM compound WHERE name = ?"
	err = db.Get(&ID, query, compoundName)
	if err != nil {
		return "", err
	}
	return ID, err
}

// Retrieves compound ID with the most similar name to the given compound name
func GetLikeCompoundID(compoundName string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var ID string
	query := "SELECT ID FROM compound WHERE name LIKE ?"
	err = db.Get(&ID, query, compoundName)
	if err != nil {
		return "", err
	}
	return ID, err
}

// Retrieves compound ID given an inchi string
func GetCompoundIDFromInchi(inchi string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var ID string
	query := "SELECT ID FROM compound WHERE inchistring = ?"
	err = db.Get(&ID, query, inchi)
	if err != nil {
		return "", err
	}
	return ID, err
}

// Retrieves inchi string given a compound ID
func GetCompoundInchi(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var inchi string
	query := "SELECT inchistring FROM compound WHERE ID = ?"
	err = db.Get(&inchi, query, compoundID)
	if err != nil {
		return "", err
	}
	return inchi, err
}

// Retrieves compound name given a compound ID
func GetCompoundName(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var name string
	query := "SELECT name FROM compound WHERE ID = ?"
	err = db.Get(&name, query, compoundID)
	if err != nil {
		return "", err
	}
	return name, err
}

// Retrieves compound name given an inchi string
func GetCompoundNameFromInchi(inchi string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var name string
	query := "SELECT name FROM compound WHERE inchistring = ?"
	err = db.Get(&name, query, inchi)
	if err != nil {
		return "", err
	}
	return name, err
}

// Retrieves the compartment that the compound is in
func GetCompoundCompartment(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var compartment string
	query := "SELECT compartment FROM compound WHERE ID = ?"
	err = db.Get(&compartment, query, compoundID)
	if err != nil {
		return "", err
	}
	return compartment, err
}

// Retrieves name of the reaction given the reaction ID
func GetReactionName(reactionID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var name string
	query := "SELECT name FROM reaction WHERE ID = ?"
	err = db.Get(&name, query, reactionID)
	if err != nil {
		return "", err
	}
	return name, err
}

// Retrieves reaction IDs that have a given compound ID as a reactant or product
func GetReactionIDsFromCompound(compoundID string, isProduct bool) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var reactionIDs []string
	query := "SELECT reaction_ID FROM reaction_compound WHERE cpd_ID = ? AND is_product = ?"
	err = db.Select(&reactionIDs, query, compoundID, isProduct)
	if err != nil {
		return nil, err
	}
	return reactionIDs, err
}

// Retrieves Model IDs that are in a given a reaction
func GetReactionSpecies(reactionID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var species []string
	query := "SELECT model_ID FROM model_reaction INDEXED BY modelreaction_ind2 WHERE reaction_ID = ?"
	err = db.Select(&species, query, reactionID)
	if err != nil {
		return nil, err
	}
	return species, err
}

// Retrieves reactions (Reaction IDs) that have a given compound (ID) as a reactant
func GetReactantCompoundIDs(reactionID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compoundIDs []string
	query := "SELECT cpd_ID FROM reaction_compound WHERE reaction_ID = ? AND is_product = ?"
	err = db.Select(&compoundIDs, query, reactionID, false)
	if err != nil {
		return nil, err
	}
	return compoundIDs, err
}

// Retrieves reactions (reaction IDs) that have a given compound (ID) as a product
func GetReactionsWithProduct(compoundID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compoundIDs []string
	query := "SELECT reaction_ID FROM reaction_compound WHERE cpd_ID = ? AND is_prod = ?"
	err = db.Select(&compoundIDs, query, compoundID, true)
	if err != nil {
		return nil, err
	}
	return compoundIDs, err
}

// Retrieves products (compound IDs) of a given reaction (ID)
func GetProductCompundIDs(reactionID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var products []string
	query := "SELECT cpd_ID FROM reaction_compound WHERE reaction_ID = ? AND is_prod = ?"
	err = db.Select(&products, query, reactionID, true)
	if err != nil {
		return nil, err
	}
	return products, err
}

// Retrives all compounds in a metabolic model given model ID
func GetModelCompounds(modelID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compounds []string
	query := "SELECT cpd_ID FROM model_compound WHERE model_ID = ?"
	err = db.Select(&compounds, query, modelID)
	if err != nil {
		return nil, err
	}
	return compounds, err
}

// Retrieves all comounds in the database
func GetAllCompoundIDs() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compounds []string
	query := "SELECT ID FROM compound"
	err = db.Select(&compounds, query)
	if err != nil {
		return nil, err
	}
	return compounds, err
}

// Retrieve all compounds in the database
func GetAllCompounds() []Compound {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	var compounds []Compound
	query := "SELECT * FROM compound"
	err = db.Select(&compounds, query)
	if err != nil {
		panic(err)
	}
	return compounds
}

// Retrieves all compound inchistrings in the database
func GetAllCompoundInchistrings() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var inchistrings []string
	query := "SELECT inchistring FROM compound"
	err = db.Select(&inchistrings, query)
	if err != nil {
		return nil, err
	}
	return inchistrings, err
}

// Retrieves all reactions in a metabolic model given model ID
func GetModelReactions(modelID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var reactions []string
	query := "SELECT reaction_ID FROM model_reaction WHERE model_ID = ?"
	err = db.Select(&reactions, query, modelID)
	if err != nil {
		return nil, err
	}
	return reactions, err
}

// Retrieves all reactions in the database
func GetAllReactions() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var reactions []string
	query := "SELECT ID FROM reaction"
	err = db.Select(&reactions, query)
	if err != nil {
		return nil, err
	}
	return reactions, err
}

// Retrieves reverisbility information of a reaction in a specified metabolic model (model ID)
func GetReactionReversibility(reactionID string, modelID string) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, err
	}
	var reversible bool
	query := "SELECT is_rev FROM model_reaction WHERE reaction_ID = ? AND model_ID = ?"
	err = db.Get(&reversible, query, reactionID, modelID)
	if err != nil {
		return false, err
	}
	return reversible, err
}

// Retrieves reversibility information of a reaction independent of model
func GetReactionReversibilityGlobal(reactionID string) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, err
	}
	var reversible bool
	query := "SELECT is_reversible FROM reaction_reversibility WHERE reaction_ID = ?"
	err = db.Get(&reversible, query, reactionID)
	if err != nil {
		return false, err
	}
	return reversible, err
}

// Retrieves gene associations for a reaction of a given metabolic network (model ID)
func GetReactionGeneAssociations(reactionID string, modelID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var genes []string
	query := "SELECT gene_ID FROM reaction_gene WHERE reaction_ID = ? AND model_ID = ?"
	err = db.Select(&genes, query, reactionID, modelID)
	if err != nil {
		return nil, err
	}
	return genes, err
}

// Retrieves protein associations for a reaction of a given metabolic network (model ID)
func GetReactionProteinAssociations(reactionID string, modelID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var proteins []string
	query := "SELECT protein_ID FROM reaction_protein WHERE reaction_ID = ? AND model_ID = ?"
	err = db.Select(&proteins, query, reactionID, modelID)
	if err != nil {
		return nil, err
	}
	return proteins, err
}

// Retrieves stoichiometry of a compound for a given reaction
func GetStoichiometry(reactionID string, compoundID string, isProduct bool) (float64, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	var stoichiometry float64
	query := "SELECT stoichiometry FROM reaction_compound WHERE reaction_ID = ? AND cpd_ID = ? AND is_prod = ?"
	err = db.Get(&stoichiometry, query, reactionID, compoundID, isProduct)
	if err != nil {
		return 0, err
	}
	return stoichiometry, err
}

// Retrieves the catalyst of reaction
func GetReactionCatalysts(reactionID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var catalysts []string
	query := "SELECT catalysts_ID FROM reaction_catalysts WHERE reaction_ID = ?"
	err = db.Get(&catalysts, query, reactionID)
	if err != nil {
		return nil, err
	}
	return catalysts, err
}

// Retrieves the compartment ID
func GetCompartmentID(compartmentName string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var compartmentID string
	query := "SELECT ID FROM compartments WHERE name LIKE ?"
	err = db.Get(&compartmentID, query, compartmentName)
	if err != nil {
		return "", err
	}
	return compartmentID, err
}

// Retrieves solvents of reaction
func GetReactionSolvents(reactionID string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var solvent []string
	query := "SELECT solvents_ID FROM reaction_solvents WHERE reaction_ID = ?"
	err = db.Get(&solvent, query, reactionID)
	if err != nil {
		return nil, err
	}
	return solvent, err
}

// Retrieves the temperaature of the reaction is performed at
func GetReactionTemperature(reactionID string) (float64, error) {
	db, err := ConnectDB()
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	var temperature float64
	query := "SELECT temperature FROM reaction_spresi_info WHERE reaction_ID = ?"
	err = db.Get(&temperature, query, reactionID)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	return temperature, err
}

// Retrieves the pressure reaction is performed at
func GetReactionPressure(reactionID string) (float64, error) {
	db, err := ConnectDB()
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	var pressure float64
	query := "SELECT pressure FROM reaction_spresi_info WHERE reaction_ID = ?"
	err = db.Get(&pressure, query, reactionID)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	return pressure, err
}

// Retrieves the time that is required to perform reaction
func GetReactionTime(reactionID string) (float64, error) {
	db, err := ConnectDB()
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	var time float64
	query := "SELECT total_time FROM reaction_spresi_info WHERE reaction_ID = ?"
	err = db.Get(&time, query, reactionID)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	return time, err
}

// Retrieves yield that was reported with reaction
func GetReactionYield(reactionID string) (float64, error) {
	db, err := ConnectDB()
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	var yield float64
	query := "SELECT yield FROM reaction_spresi_info WHERE reaction_ID = ?"
	err = db.Get(&yield, query, reactionID)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	return yield, err
}

// Retrieves the reference of reaction
func GetReactionReference(reactionID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var reference string
	query := "SELECT reference FROM reaction_spresi_info WHERE reaction_ID = ?"
	err = db.Get(&reference, query, reactionID)
	if err != nil {
		return "", err
	}
	return reference, err
}

// Retrieves reactions based on type
func GetReactionsByType(reactionType string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var reactions []string
	query := "SELECT ID FROM reactions WHERE type = ?"
	err = db.Select(&reactions, query, reactionType)
	if err != nil {
		return nil, err
	}
	return reactions, err
}

// Retrieves reaction on type
func GetReactionType(reactionID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var reactionType string
	query := "SELECT type FROM reactions WHERE ID = ?"
	err = db.Get(&reactionType, query, reactionID)
	if err != nil {
		return "", err
	}
	return reactionType, err
}

// Retrieves all reaction KEGG IDs
func GetAllReactionKEGGIDs() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var reactionIDs []string
	query := "SELECT kegg_id FROM reaction"
	err = db.Select(&reactionIDs, query)
	if err != nil {
		return nil, err
	}
	return reactionIDs, err
}

// Retrieves kegg ID for a reaction based on main ID
func GetReactionKEGGID(reactionID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var reactionKEGGID string
	query := "SELECT kegg_id FROM reaction WHERE ID = ?"
	err = db.Get(&reactionKEGGID, query, reactionID)
	if err != nil {
		return "", err
	}
	return reactionKEGGID, err
}

// Retrieves kegg ID for a compound based on main ID
func GetCompoundKEGGID(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var compoundKEGGID string
	query := "SELECT kegg_id FROM compound WHERE ID = ?"
	err = db.Get(&compoundKEGGID, query, compoundID)
	if err != nil {
		return "", err
	}
	return compoundKEGGID, err
}

// Retrieves all compound Kegg IDs
func GetAllCompoundKEGGIDs() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compoundIDs []string
	query := "SELECT kegg_id FROM compound"
	err = db.Select(&compoundIDs, query)
	if err != nil {
		return nil, err
	}
	return compoundIDs, err
}

// Retrieves all chemicalformulas
func GetAllChemicalFormulas() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var formulas []string
	query := "SELECT chemicalformula FROM compound"
	err = db.Select(&formulas, query)
	if err != nil {
		return nil, err
	}
	return formulas, err
}

// Retrieves chemicalformula for compound ID
func GetChemicalFormula(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var formula string
	query := "SELECT chemicalformula FROM compound WHERE ID = ?"
	err = db.Get(&formula, query, compoundID)
	if err != nil {
		return "", err
	}
	return formula, err
}

// Retrieves casnumber for compound ID
func GetCASNumber(compoundID string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var casNumber string
	query := "SELECT casnumber FROM compound WHERE ID = ?"
	err = db.Get(&casNumber, query, compoundID)
	if err != nil {
		return "", err
	}
	return casNumber, err
}

// Retrieves compound IDs for Chemical formula
func GetCompoundIDByFormula(formula string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compoundIDs []string
	query := "SELECT ID FROM compound WHERE chemicalformula = ?"
	err = db.Get(&compoundIDs, query, formula)
	if err != nil {
		return nil, err
	}
	return compoundIDs, err
}

// Retrieves compound name for given search term (name/formula) TODO: Add more match criteria
func GetCompoundNameBySearchTerm(searchTerm string) ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var compoundNames []string
	query := "SELECT name FROM compound WHERE name LIKE ? OR chemicalformula LIKE ? OR ID LIKE ? OR kegg_id LIKE ? OR casnumber LIKE ?"
	err = db.Get(&compoundNames, query, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, err
	}
	return compoundNames, err
}

// Retrieves model ID for given file_name
func GetModelIDByFileName(fileName string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	var modelID string
	query := "SELECT ID FROM model WHERE file_name = ?"
	err = db.Get(&modelID, query, fileName)
	if err != nil {
		return "", err
	}
	return modelID, err
}

// Retrieves all model IDs in the database
func GetAllFBAModelIDs() ([]string, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	var modelIDs []string
	query := "SELECT ID FROM fba_models"
	err = db.Select(&modelIDs, query)
	if err != nil {
		return nil, err
	}
	return modelIDs, err
}
