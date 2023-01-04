/*
Package retsynth_test contains tests for the queries.go in the retsynth package.
*/
package retsynth

import (
	"testing"
)

//	func setupSuite() {
//		// Setup the path for the OS environment variable for the Retsynth database
//	}
func TestConnectDB(t *testing.T) {
	_, err := ConnectDB()
	if err != nil {
		t.Error("Error connecting to the database")
	}
}

func TestGetUniqueMetabolicClusters(t *testing.T) {
	_, err := GetUniqueMetabolicClusters()
	if err != nil {
		t.Error("Error getting unique metabolic clusters")
	}
}

func TestGetAllModelIDs(t *testing.T) {
	_, err := GetAllModelIDs()
	if err != nil {
		t.Error("Error getting all model ids")
	}
}

func TestGetOrganismName(t *testing.T) {
	var name string = "Escherichia coli"
	_, err := GetOrganismName(name)
	if err != nil {
		t.Error("Error getting organism names")
	}
}

func TestGetOrganismID(t *testing.T) {
	var organismID string = "83333"
	_, err := GetOrganismID(organismID)
	if err != nil {
		t.Error("Error getting organism ids")
	}
}

func TestGetCompoundID(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCompoundID(compoundID)
	if err != nil {
		t.Error("Error getting compound ids")
	}
}

func TestGetLikeCompoundID(t *testing.T) {
	var compundName string = "glucose"
	_, err := GetLikeCompoundID(compundName)
	if err != nil {
		t.Error("Error getting like compound ids")
	}
}

func TestGetCompoundIDFromInchi(t *testing.T) {
	var inchistring string = "InChI=1S/C3H4O3S/c4-2(1-7)3(5)6/h7H,1H2,(H,5,6)_c0"
	_, err := GetCompoundIDFromInchi(inchistring)
	if err != nil {
		t.Error("Error getting compound ids from inchi")
	}
}

func TestGetCompoundInchi(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCompoundInchi(compoundID)
	if err != nil {
		t.Error("Error getting compound inchi")
	}
}

func TestGetCompoundName(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCompoundName(compoundID)
	if err != nil {
		t.Error("Error getting compound names")
	}
}

func TestGetCompoundNameFromInchi(t *testing.T) {
	var inchistring string = "InChI=1S/C5H11O8P/c6-1-3(7)5(9)4(8)2-13-14(10,11)12/h1,3-5,7-9H,2H2,(H2,10,11,12)/t3-,4-,5+/m1/s1_c0"
	_, err := GetCompoundNameFromInchi(inchistring)
	if err != nil {
		t.Error("Error getting compound names from inchi")
	}
}

func TestGetCompoundCompartment(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCompoundCompartment(compoundID)
	if err != nil {
		t.Error("Error getting compound compartments")
	}
}

func TestGetReactionName(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionName(reactionID)
	if err != nil {
		t.Error("Error getting reaction names")
	}
}

func TestGetReactionIDsFromCompound(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetReactionIDsFromCompound(compoundID, true)
	if err != nil {
		t.Error("Error getting reaction ids from compound")
	}

	// Test for false | Figure out the exact assertions and the test case
	_, err = GetReactionIDsFromCompound(compoundID, false)
	if err != nil {
		t.Error("Error getting reaction ids from compound")
	}

}

func TestGetReactionSpecies(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionSpecies(reactionID)
	if err != nil {
		t.Error("Error getting reaction species")
	}
}

func TestGetReactantCompoundIDs(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactantCompoundIDs(reactionID)
	if err != nil {
		t.Error("Error getting reactant compound ids")
	}
}

func TestGetReactionsWithProduct(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetReactionsWithProduct(compoundID)
	if err != nil {
		t.Error("Error getting reactions with product")
	}
}

func TestGetProductCompundIDs(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetProductCompundIDs(reactionID)
	if err != nil {
		t.Error("Error getting product compound ids")
	}
}

func TestGetModelCompounds(t *testing.T) {
	var modelID string = "iJO1366"
	_, err := GetModelCompounds(modelID)
	if err != nil {
		t.Error("Error getting model compounds")
	}
}

func TestGetAllCompounds(t *testing.T) {
	_, err := GetAllCompoundIDs()
	if err != nil {
		t.Error("Error getting all compounds")
	}
}

func TestGetAllCompoundInchistrings(t *testing.T) {
	_, err := GetAllCompoundInchistrings()
	if err != nil {
		t.Error("Error getting all compound inchistrings")
	}
}

func TestGetModelReactions(t *testing.T) {
	var modelID string = "iJO1366"
	_, err := GetModelReactions(modelID)
	if err != nil {
		t.Error("Error getting model reactions")
	}
}

func TestGetAllReactions(t *testing.T) {
	_, err := GetAllReactions()
	if err != nil {
		t.Error("Error getting all reactions")
	}
}

func TestGetReactionReversibility(t *testing.T) {
	var reactionID string = "rxn00001"
	var modelID string = "iJO1366"
	_, err := GetReactionReversibility(reactionID, modelID)
	if err != nil {
		t.Error("Error getting reaction reversibility")
	}
}

func TestGetReactionReversibilityGlobal(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionReversibilityGlobal(reactionID)
	if err != nil {
		t.Error("Error getting reaction reversibility global")
	}
}

func TestGetReactionGeneAssociations(t *testing.T) {
	var reactionID string = "rxn00001"
	var modelID string = "iJO1366"
	_, err := GetReactionGeneAssociations(reactionID, modelID)
	if err != nil {
		t.Error("Error getting reaction gene associations")
	}
}

func TestGetReactionProteinAssociations(t *testing.T) {
	var reactionID string = "rxn00001"
	var modelID string = "iJO1366"
	_, err := GetReactionProteinAssociations(reactionID, modelID)
	if err != nil {
		t.Error("Error getting reaction protein associations")
	}
}

func TestGetStoichiometry(t *testing.T) {
	var reactionID string = "rxn00001"
	var compoundID string = "cpd00001"
	var isProduct bool = true
	_, err := GetStoichiometry(reactionID, compoundID, isProduct)
	if err != nil {
		t.Error("Error getting stoichiometry")
	}
}

func TestGetReactionCatalyst(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionCatalysts(reactionID)
	if err != nil {
		t.Error("Error getting reaction catalysts")
	}
}

func TestGetCompartmentID(t *testing.T) {
	var compartmentID string = "c"
	_, err := GetCompartmentID(compartmentID)
	if err != nil {
		t.Error("Error getting compartment ids")
	}
}

func TestGetReactionSolvents(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionSolvents(reactionID)
	if err != nil {
		t.Error("Error getting reaction solvents")
	}
}

func TestGetReactionTemperature(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionTemperature(reactionID)
	if err != nil {
		t.Error("Error getting reaction temperature")
	}
}

func TestGetReactionPressure(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionPressure(reactionID)
	if err != nil {
		t.Error("Error getting reaction pressure")
	}
}

func TestGetReactionTime(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionTime(reactionID)
	if err != nil {
		t.Error("Error getting reaction time")
	}
}

func TestGetReactionYield(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionYield(reactionID)
	if err != nil {
		t.Error("Error getting reaction yield")
	}
}

func TestGetReactionReference(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionReference(reactionID)
	if err != nil {
		t.Error("Error getting reaction references")
	}
}

func TestGetReactionsByType(t *testing.T) {
	var reactionType string = "bio"
	_, err := GetReactionsByType(reactionType)
	if err != nil {
		t.Error("Error getting reactions by type")
	}
}

func TestGetReactionType(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionType(reactionID)
	if err != nil {
		t.Error("Error getting reaction type")
	}
}

func TestGetAllReactionKEGGIDs(t *testing.T) {
	_, err := GetAllReactionKEGGIDs()
	if err != nil {
		t.Error("Error getting all reaction kegg ids")
	}
}

func TestGetReactionKEGGID(t *testing.T) {
	var reactionID string = "rxn00001"
	_, err := GetReactionKEGGID(reactionID)
	if err != nil {
		t.Error("Error getting reaction kegg id")
	}
}

func TestGetCompoundKEGGID(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCompoundKEGGID(compoundID)
	if err != nil {
		t.Error("Error getting compound kegg id")
	}
}

func TestGetAllCompoundKEGGIDs(t *testing.T) {
	_, err := GetAllCompoundKEGGIDs()
	if err != nil {
		t.Error("Error getting all compound kegg ids")
	}
}

func TestGetAllChemicalFormulas(t *testing.T) {
	_, err := GetAllChemicalFormulas()
	if err != nil {
		t.Error("Error getting all chemical formulas")
	}
}

func TestGetChemicalFormula(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetChemicalFormula(compoundID)
	if err != nil {
		t.Error("Error getting chemical formula")
	}
}

func TestGetCASNumber(t *testing.T) {
	var compoundID string = "cpd00001"
	_, err := GetCASNumber(compoundID)
	if err != nil {
		t.Error("Error getting cas number")
	}
}

func TestGetCompoundIDByFormula(t *testing.T) {
	var formula string = "C6H12O6"
	_, err := GetCompoundIDByFormula(formula)
	if err != nil {
		t.Error("Error getting compound id by formula")
	}
}

func TestGetCompoundNameBySearchTerm(t *testing.T) {
	var searchTerm string = "glucose"
	_, err := GetCompoundNameBySearchTerm(searchTerm)
	if err != nil {
		t.Error("Error getting compound name by search term")
	}
}

func TestGetModelIDByFileName(t *testing.T) {
	var fileName string = "iJO1366.xml"
	_, err := GetModelIDByFileName(fileName)
	if err != nil {
		t.Error("Error getting model id by file name")
	}
}

func TestGetAllFBAModelIDs(t *testing.T) {
	_, err := GetAllFBAModelIDs()
	if err != nil {
		t.Error("Error getting all fba model ids")
	}
}
