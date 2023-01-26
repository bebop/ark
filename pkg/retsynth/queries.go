package retsynth

import (
	"database/sql"

	"github.com/TimothyStiles/allbase/parameters"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Easy database connector
func ConnectDB() *sqlx.DB {
	var db, err = sqlx.Connect("sqlite3", parameters.RetsynthDBPath())
	if err != nil {
		panic(err)
	}
	return db
}

// Retrieves unique metabolic clusters (organisms with the exact same metabolism) in the database
func GetUniqueMetabolicClusters() []string {
	db := ConnectDB()
	var clusters []string
	query := "SELECT DISTINCT cluster_num FROM cluster"
	var err = db.Select(&clusters, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return clusters
}

// Retrieves model IDs from a specified cluster in the database
func GetModelIDsFromCluster(cluster string) []string {
	db := ConnectDB()
	var models []string
	query := "SELECT ID FROM cluster WHERE cluster_num = ?"
	var err = db.Select(&models, query, cluster)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return models
}

// Retrieves all model IDs from the database
func GetAllModelIDs() []string {
	db := ConnectDB()
	var models []string
	query := "SELECT DISTINCT ID FROM model"
	var err = db.Select(&models, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return models
}

// Retrieves all model objects from the database
func GetAllModels() []Model {
	db := ConnectDB()
	var models []Model
	query := "SELECT * FROM model"
	var err = db.Select(&models, query)
	if err != nil {
		panic(err)
	}
	return models
}

// Retrieves name of organism given a specific organism ID reutrns empty string if not found
func GetOrganismName(organismID string) sql.NullString {
	db := ConnectDB()
	var name string
	query := "SELECT file_name FROM model WHERE ID = ?"
	var err = db.Get(&name, query, organismID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: name, Valid: true}
}

// Retrieves ID of organism given a specific organism name
func GetOrganismID(organismName string) sql.NullString {
	db := ConnectDB()
	var ID string
	query := "SELECT ID FROM model WHERE file_name = ?"
	var err = db.Get(&ID, query, organismName)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: ID, Valid: true}
}

// Retrieves compound ID given a compound name
func GetCompoundID(compoundName string) sql.NullString {
	db := ConnectDB()
	var ID string
	query := "SELECT ID FROM compound WHERE name = ?"
	var err = db.Get(&ID, query, compoundName)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: ID, Valid: true}
}

// Retrieves compound ID with the most similar name to the given compound name
func GetLikeCompoundID(compoundName string) sql.NullString {
	db := ConnectDB()
	var ID string
	query := "SELECT ID FROM compound WHERE name LIKE ?"
	var err = db.Get(&ID, query, compoundName)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: ID, Valid: true}
}

// Retrieves compound ID given an inchi string
func GetCompoundIDFromInchi(inchi string) sql.NullString {
	db := ConnectDB()
	var ID string
	query := "SELECT ID FROM compound WHERE inchistring = ?"
	var err = db.Get(&ID, query, inchi)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: ID, Valid: true}
}

// Retrieves inchi string given a compound ID
func GetCompoundInchi(compoundID string) sql.NullString {
	db := ConnectDB()
	var inchi string
	query := "SELECT inchistring FROM compound WHERE ID = ?"
	var err = db.Get(&inchi, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: inchi, Valid: true}
}

// Retrieves compound name given a compound ID
func GetCompoundName(compoundID string) sql.NullString {
	db := ConnectDB()
	var name string
	query := "SELECT name FROM compound WHERE ID = ?"
	var err = db.Get(&name, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: name, Valid: true}
}

// Retrieves compound name given an inchi string
func GetCompoundNameFromInchi(inchi string) sql.NullString {
	db := ConnectDB()
	var name string
	query := "SELECT name FROM compound WHERE inchistring = ?"
	var err = db.Get(&name, query, inchi)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: name, Valid: true}
}

// Retrieves the compartment that the compound is in
func GetCompoundCompartment(compoundID string) sql.NullString {
	db := ConnectDB()
	var compartment string
	query := "SELECT compartment FROM compound WHERE ID = ?"
	var err = db.Get(&compartment, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: compartment, Valid: true}
}

// Retrieves name of the reaction given the reaction ID
func GetReactionName(reactionID string) sql.NullString {
	db := ConnectDB()
	var name string
	query := "SELECT name FROM reaction WHERE ID = ?"
	var err = db.Get(&name, query, reactionID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: name, Valid: true}
}

// Retrieves reaction IDs that have a given compound ID as a reactant or product
func GetReactionIDsFromCompound(compoundID string, isProduct bool) []string {
	db := ConnectDB()
	var reactionIDs []string
	query := "SELECT reaction_ID FROM reaction_compound WHERE cpd_ID = ? AND is_prod = ?"
	var err = db.Select(&reactionIDs, query, compoundID, isProduct)
	if err != nil {
		panic(err)
	}
	return reactionIDs
}

// Retrieves Model IDs that are in a given a reaction
func GetReactionSpecies(reactionID string) []string {
	db := ConnectDB()
	var species []string
	query := "SELECT model_ID FROM model_reaction INDEXED BY modelreaction_ind2 WHERE reaction_ID = ?"
	var err = db.Select(&species, query, reactionID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return species
}

// Retrieves reactions (Reaction IDs) that have a given compound (ID) as a reactant
func GetReactantCompoundIDs(reactionID string) []string {
	db := ConnectDB()
	var compoundIDs []string
	query := "SELECT cpd_ID FROM reaction_compound WHERE reaction_ID = ? AND is_prod = ?"
	var err = db.Select(&compoundIDs, query, reactionID, false)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return compoundIDs
}

// Retrieves reactions (reaction IDs) that have a given compound (ID) as a product
func GetReactionsWithProduct(compoundID string) []string {
	db := ConnectDB()
	var compoundIDs []string
	query := "SELECT reaction_ID FROM reaction_compound WHERE cpd_ID = ? AND is_prod = ?"
	var err = db.Select(&compoundIDs, query, compoundID, true)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return compoundIDs
}

// Retrieves products (compound IDs) of a given reaction (ID)
func GetProductCompundIDs(reactionID string) []string {
	db := ConnectDB()
	var products []string
	query := "SELECT cpd_ID FROM reaction_compound WHERE reaction_ID = ? AND is_prod = ?"
	var err = db.Select(&products, query, reactionID, true)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return products
}

// Retrives all compounds in a metabolic model given model ID
func GetModelCompounds(modelID string) []string {
	db := ConnectDB()
	var compounds []string
	query := "SELECT cpd_ID FROM model_compound WHERE model_ID = ?"
	var err = db.Select(&compounds, query, modelID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compounds
}

// Retrieves all comounds in the database
func GetAllCompoundIDs() []string {
	db := ConnectDB()
	var compounds []string
	query := "SELECT ID FROM compound"
	var err = db.Select(&compounds, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compounds
}

// Retrieve all compounds in the database
func GetAllCompounds() []Compound {
	db := ConnectDB()
	var compounds []Compound
	query := "SELECT * FROM compound"
	var err = db.Select(&compounds, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compounds
}

// Retrieves all compound inchistrings in the database
func GetAllCompoundInchiStrings() []string {
	db := ConnectDB()
	var inchistrings []sql.NullString
	query := "SELECT inchistring FROM compound"
	var err = db.Select(&inchistrings, query)
	if err != nil {
		return nil
	}
	// Convert sql.NullString to string
	var returninchistrings []string
	for _, inchistring := range inchistrings {
		if inchistring.Valid {
			returninchistrings = append(returninchistrings, inchistring.String)
		}
	}
	return returninchistrings
}

// Retrieves all reactions in a metabolic model given model ID
func GetModelReactions(modelID string) []string {
	db := ConnectDB()
	var reactions []string
	query := "SELECT reaction_ID FROM model_reaction WHERE model_ID = ?"
	var err = db.Select(&reactions, query, modelID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return reactions
}

// Retrieves all reactions in the database
func GetAllReactions() []string {
	db := ConnectDB()
	var reactions []string
	query := "SELECT ID FROM reaction"
	var err = db.Select(&reactions, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return nil
	}
	return reactions
}

// Retrieves reverisbility information of a reaction in a specified metabolic model (model ID)
func GetReactionReversibility(reactionID string, modelID string) sql.NullBool {
	db := ConnectDB()
	var reversible bool
	query := "SELECT is_rev FROM model_reaction WHERE reaction_ID = ? AND model_ID = ? LIMIT 1"
	var err = db.Get(&reversible, query, reactionID, modelID)
	if err == sql.ErrNoRows {
		return sql.NullBool{Bool: false, Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullBool{Bool: reversible, Valid: true}
}

// Retrieves reversibility information of a reaction independent of model
func GetReactionReversibilityGlobal(reactionID string) sql.NullBool {
	db := ConnectDB()
	var reversible bool
	query := "SELECT is_reversible FROM reaction_reversibility WHERE reaction_ID = ?"
	var err = db.Get(&reversible, query, reactionID)
	if err != nil {
		return sql.NullBool{Bool: false, Valid: false}
	}
	return sql.NullBool{Bool: reversible, Valid: true}
}

// Retrieves gene associations for a reaction of a given metabolic network (model ID)
func GetReactionGeneAssociations(reactionID string, modelID string) []string {
	db := ConnectDB()
	var genes []string
	query := "SELECT gene_ID FROM reaction_gene WHERE reaction_ID = ? AND model_ID = ?"
	var err = db.Select(&genes, query, reactionID, modelID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return genes
}

// Retrieves protein associations for a reaction of a given metabolic network (model ID)
func GetReactionProteinAssociations(reactionID string, modelID string) []string {
	db := ConnectDB()
	var proteins []string
	query := "SELECT protein_ID FROM reaction_protein WHERE reaction_ID = ? AND model_ID = ?"
	var err = db.Select(&proteins, query, reactionID, modelID)
	if err != nil {
		return nil
	}
	return proteins
}

// Retrieves stoichiometry of a compound for a given reaction
func GetStoichiometry(reactionID string, compoundID string, isProduct bool) sql.NullFloat64 {
	db := ConnectDB()
	var stoichiometry float64
	query := "SELECT stoichiometry FROM reaction_compound WHERE reaction_ID = ? AND cpd_ID = ? AND is_prod = ? LIMIT 1"
	var err = db.Get(&stoichiometry, query, reactionID, compoundID, isProduct)
	if err != nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: stoichiometry, Valid: true}
}

// Retrieves the catalyst of reaction
func GetReactionCatalysts(reactionID string) []string {
	//TODO: #64 The current schema doesnt have catalyst information, so this function fails
	db := ConnectDB()
	var catalysts []string
	query := "SELECT catalysts_ID FROM reaction_catalysts WHERE reaction_ID = ?"
	var err = db.Select(&catalysts, query, reactionID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return catalysts
}

// Retrieves the compartment ID
func GetCompartmentID(compartmentName string) sql.NullString {
	db := ConnectDB()
	var compartmentID string
	query := "SELECT ID FROM compartments WHERE name = ?"
	var err = db.Get(&compartmentID, query, compartmentName)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: compartmentID, Valid: true}
}

// Retrieves solvents of reaction
func GetReactionSolvents(reactionID string) []string {
	// TODO: The current schema doesnt have solvent information, so this function fails
	db := ConnectDB()
	var solvent []string
	query := "SELECT solvents_ID FROM reaction_solvents WHERE reaction_ID = ?"
	var err = db.Select(&solvent, query, reactionID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return solvent
}

// Retrieves the temperaature of the reaction is performed at
func GetReactionTemperature(reactionID string) sql.NullFloat64 {
	db := ConnectDB()
	var temperature float64
	query := "SELECT temperature FROM reaction_spresi_info WHERE reaction_ID = ?"
	var err = db.Get(&temperature, query, reactionID)
	if err != nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: temperature, Valid: true}
}

// Retrieves the pressure reaction is performed at
func GetReactionPressure(reactionID string) sql.NullFloat64 {
	db := ConnectDB()
	var pressure float64
	query := "SELECT pressure FROM reaction_spresi_info WHERE reaction_ID = ?"
	var err = db.Get(&pressure, query, reactionID)
	if err != nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: pressure, Valid: true}
}

// Retrieves the time that is required to perform reaction
func GetReactionTime(reactionID string) sql.NullFloat64 {
	db := ConnectDB()
	var time float64
	query := "SELECT total_time FROM reaction_spresi_info WHERE reaction_ID = ?"
	var err = db.Get(&time, query, reactionID)
	if err != nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: time, Valid: true}
}

// Retrieves yield that was reported with reaction
func GetReactionYield(reactionID string) sql.NullFloat64 {
	db := ConnectDB()
	var yield float64
	query := "SELECT yield FROM reaction_spresi_info WHERE reaction_ID = ?"
	var err = db.Get(&yield, query, reactionID)
	if err != nil {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: yield, Valid: true}
}

// Retrieves the reference of reaction
func GetReactionReference(reactionID string) sql.NullString {
	db := ConnectDB()
	var reference string
	query := "SELECT reference FROM reaction_spresi_info WHERE reaction_ID = ?"
	var err = db.Get(&reference, query, reactionID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: reference, Valid: true}
}

// Retrieves reactions based on type
func GetReactionsByType(reactionType string) []string {
	db := ConnectDB()
	var reactions []string
	query := "SELECT ID FROM reaction WHERE type = ?"
	var err = db.Select(&reactions, query, reactionType)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return reactions
}

// Retrieves reaction on type
func GetReactionType(reactionID string) sql.NullString {
	db := ConnectDB()
	var reactionType string
	query := "SELECT type FROM reaction WHERE ID = ?"
	var err = db.Get(&reactionType, query, reactionID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: reactionType, Valid: true}
}

// Retrieves all reaction KEGG IDs
func GetAllReactionKEGGIDs() []string {
	db := ConnectDB()
	var reactionIDs []sql.NullString
	query := "SELECT kegg_id FROM reaction"
	var err = db.Select(&reactionIDs, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	// Convert sql.NullString to string
	var reactionKEGGIDs []string
	for _, reactionID := range reactionIDs {
		if reactionID.Valid {
			reactionKEGGIDs = append(reactionKEGGIDs, reactionID.String)
		}
	}
	return reactionKEGGIDs
}

// Retrieves kegg ID for a reaction based on main ID
func GetReactionKEGGID(reactionID string) sql.NullString {
	db := ConnectDB()
	var reactionKEGGID string
	query := "SELECT kegg_id FROM reaction WHERE ID = ?"
	var err = db.Get(&reactionKEGGID, query, reactionID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: reactionKEGGID, Valid: true}
}

// Retrieves kegg ID for a compound based on main ID
func GetCompoundKEGGID(compoundID string) sql.NullString {
	db := ConnectDB()
	var compoundKEGGID string
	query := "SELECT kegg_id FROM compound WHERE ID = ?"
	var err = db.Get(&compoundKEGGID, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: compoundKEGGID, Valid: true}
}

// Retrieves all compound Kegg IDs
func GetAllCompoundKEGGIDs() []string {
	db := ConnectDB()
	var compoundIDs []string
	query := "SELECT kegg_id FROM compound"
	var err = db.Select(&compoundIDs, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compoundIDs
}

// Retrieves all chemicalformulas
func GetAllChemicalFormulas() []string {
	db := ConnectDB()
	var formulas []sql.NullString
	query := "SELECT chemicalformula FROM compound"
	var err = db.Select(&formulas, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	// Convert sql.NullString to string
	var returnFormulas []string
	for _, formula := range formulas {
		if formula.Valid {
			returnFormulas = append(returnFormulas, formula.String)
		}
	}
	return returnFormulas
}

// Retrieves chemicalformula for compound ID
func GetChemicalFormula(compoundID string) sql.NullString {
	db := ConnectDB()
	var formula string
	query := "SELECT chemicalformula FROM compound WHERE ID = ?"
	var err = db.Get(&formula, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: formula, Valid: true}
}

// Retrieves casnumber for compound ID
func GetCASNumber(compoundID string) sql.NullString {
	db := ConnectDB()
	var casNumber string
	query := "SELECT casnumber FROM compound WHERE ID = ?"
	var err = db.Get(&casNumber, query, compoundID)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: casNumber, Valid: true}
}

// Retrieves compound IDs for Chemical formula
func GetCompoundIDByFormula(formula string) []string {
	db := ConnectDB()
	var compoundIDs []string
	query := "SELECT ID FROM compound WHERE chemicalformula = ?"
	var err = db.Select(&compoundIDs, query, formula)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compoundIDs
}

// Retrieves compound name for given search term (name/formula) TODO: Add more match criteria
func GetCompoundBySearchTerm(searchTerm string) []Compound {
	// Update the search term with the % wildcard
	searchTerm = "%" + searchTerm + "%"
	db := ConnectDB()
	var compoundNames []Compound
	query := "SELECT * FROM compound WHERE name LIKE ? OR chemicalformula LIKE ? OR ID LIKE ? OR kegg_id LIKE ? OR casnumber LIKE ? OR inchistring LIKE ?"
	var err = db.Select(&compoundNames, query, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return compoundNames
}

// Retrieves organism name for given search term (name/formula) TODO: Add more match criteria as we go along
func GetOrganismBySearchTerm(searchTerm string) []Model {
	searchTerm = "%" + searchTerm + "%"
	db := ConnectDB()
	var organisms []Model
	query := "SELECT * FROM model WHERE file_name LIKE ? OR ID LIKE ?"
	var err = db.Select(&organisms, query, searchTerm, searchTerm)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return organisms
}

// Retrieves model ID for given file_name
func GetModelIDByFileName(fileName string) sql.NullString {
	db := ConnectDB()
	var modelID string
	query := "SELECT ID FROM model WHERE file_name = ?"
	var err = db.Get(&modelID, query, fileName)
	if err == sql.ErrNoRows {
		return sql.NullString{String: "", Valid: false}
	}
	if err != nil {
		panic(err)
	}
	return sql.NullString{String: modelID, Valid: true}
}

// Retrieves all model IDs in the database
func GetAllFBAModelIDs() []string {
	db := ConnectDB()
	var modelIDs []string
	query := "SELECT ID FROM fba_models"
	var err = db.Select(&modelIDs, query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return modelIDs
}
