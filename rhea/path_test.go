package rhea

import (
	"fmt"
	"testing"
)

func TestRheaPath_pathway(t *testing.T) {
	// We are going to target part of the formate pathway as represented
	// https://github.com/Koeng101/toolkits/blob/main/formate_challenge/doc.md
	// Specifically, we want to degrade formate into CHEBI:15636, or
	// (6R)-5,10-methylenetetrahydrofolate(2−)
	targetCompound := "http://purl.obolibrary.org/obo/CHEBI:15636"
	// Roughly speaking, this will be a 3 step pathway:
	// 6.3.4.3 -> 3.5.4.9 -> 1.5.1.5
	// We will provide all inputs to thse pathways for the production of
	// CHEBI:15636

	// 6.3.4.3 inputs
	tetrahydrofolate := "http://purl.obolibrary.org/obo/CHEBI:57453"
	atp := "http://purl.obolibrary.org/obo/CHEBI:30616"
	formate := "http://purl.obolibrary.org/obo/CHEBI:15740"

	// 3.5.4.9 inputs
	h := "http://purl.obolibrary.org/obo/CHEBI:15378"

	// 1.5.1.5 inputs
	nadph := "http://purl.obolibrary.org/obo/CHEBI:57783"

	reactionInputs := []string{tetrahydrofolate, atp, formate, h, nadph}

	// Generate RheaPath
	rheaPath := RheaToPath(rhea)
	_, _ := rheaPath.Pathways(targetCompound, reactionInputs)
}
