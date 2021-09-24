package rhea

import (
	"sync"
)

// Convert rhea into a more suitable data format.

type RheaPath struct {
	ReactionsToProductCompounds   map[string]map[string]bool
	ReactionsToSubstrateCompounds map[string]map[string]bool
	SubstrateCompoundsToReactions map[string]map[string]bool
	ProductCompoundsToReactions   map[string]map[string]bool
}

func RheaToPath(rhea Rhea) RheaPath {
	ReactionsToProductCompounds := make(map[string]map[string]bool)
	ReactionsToSubstrateCompounds := make(map[string]map[string]bool)
	SubstrateCompoundsToReactions := make(map[string]map[string]bool)
	ProductCompoundsToReactions := make(map[string]map[string]bool)

	// Establish possible reactions
	for _, reaction := range rhea.Reactions {
		ReactionsToProductCompounds[reaction.Accession] = make(map[string]bool)
		ReactionsToSubstrateCompounds[reaction.Accession] = make(map[string]bool)
	}

	// Establish possible compounds
	compoundToChebi := make(map[string]string)
	for _, compound := range rhea.Compounds {
		SubstrateCompoundsToReactions[compound.ChEBI] = make(map[string]bool)
		ProductCompoundsToReactions[compound.ChEBI] = make(map[string]bool)

		compoundToChebi[compound.Accession] = compound.ChEBI
	}

	// Substrates and Products are represented by reactionside. We can relate
	// reactionsides to compounds by making a map with ReactionParticipant.
	reactionSideToCompound := make(map[string]string)
	for _, reactionParticipant := range rhea.ReactionParticipants {
		reactionSideToCompound[reactionParticipant.ReactionSide] = compoundToChebi[reactionParticipant.Compound]
	}

	// For every reaction, establish the possible ReactionsToProductCompounds
	// and ReactionsToSubstrateCompounds.
	for _, reaction := range rhea.Reactions {
		for _, substrate := range reaction.Substrates {
			ReactionsToSubstrateCompounds[reaction.Accession][reactionSideToCompound[substrate]] = true
			SubstrateCompoundsToReactions[reactionSideToCompound[substrate]][reaction.Accession] = true
		}
		for _, product := range reaction.Products {
			ReactionsToProductCompounds[reaction.Accession][reactionSideToCompound[product]] = true
			ProductCompoundsToReactions[reactionSideToCompound[product]][reaction.Accession] = true
		}
		for _, productOrSubstrate := range reaction.SubstrateOrProducts {
			ReactionsToSubstrateCompounds[reaction.Accession][reactionSideToCompound[productOrSubstrate]] = true
			ReactionsToProductCompounds[reaction.Accession][reactionSideToCompound[productOrSubstrate]] = true

			SubstrateCompoundsToReactions[reactionSideToCompound[productOrSubstrate]][reaction.Accession] = true
			ProductCompoundsToReactions[reactionSideToCompound[productOrSubstrate]][reaction.Accession] = true
		}
	}

	return RheaPath{ReactionsToProductCompounds: ReactionsToProductCompounds, ReactionsToSubstrateCompounds: ReactionsToSubstrateCompounds, SubstrateCompoundsToReactions: SubstrateCompoundsToReactions, ProductCompoundsToReactions: ProductCompoundsToReactions}
}

func (rp *RheaPath) Pathways(targetCompound string, substrateInputs []string) ([][]string, error) {
	// maxDepth DOES.NOT.SCALE.
	maxDepth := 1
	currentDepth := 0

	// Next, establish some channels for us to work with
	// as well as a waitgroup
	var wg sync.WaitGroup
	validPathways := make(chan []string)
	outputPathways := make(chan [][]string)

	// Begin recursion
	wg.Add(1)
	go recurseReactionTree(rp, &wg, validPathways, targetCompound, substrateInputs, []string{}, currentDepth, maxDepth)
	go getPathways(validPathways, outputPathways)
	wg.Wait()
	close(validPathways)
	output := <-outputPathways
	return output, nil

}

func (rp *RheaPath) getSatisfyingReactions(substrateInputs []string, currentReactions []string) []string {
	// Substrate map
	substrateInputMap := make(map[string]bool)
	for _, reactionInput := range substrateInputs {
		substrateInputMap[reactionInput] = true
	}

	// currentReactions map
	currentReactionMap := make(map[string]bool)
	for _, currentReaction := range currentReactions {
		currentReactionMap[currentReaction] = true
	}

	var validPathway []string
	var pathwaySatisfied bool
	for reaction, substrateMap := range rp.ReactionsToSubstrateCompounds {
		pathwaySatisfied = true // we assume this is true until we find a substrate that is not included.
		// iterate through the substrates for the pathway. If all reactionInputs are satisfied, we can begin
		// with this pathway
		for substrate, _ := range substrateMap {
			_, ok := substrateInputMap[substrate]
			if !ok {
				pathwaySatisfied = false
			}
		}
		if pathwaySatisfied {
			_, ok := currentReactionMap[reaction]
			if !ok {
				validPathway = append(validPathway, reaction)
			}
		}
	}
	return validPathway
}

func recurseReactionTree(rp *RheaPath, wg *sync.WaitGroup, validPathways chan []string, targetCompound string, substrateInputs []string, currentReactions []string, currentDepth int, maxDepth int) {
	defer wg.Done()

	// Are we done here?
	for _, substrate := range substrateInputs {
		if substrate == targetCompound {
			validPathways <- currentReactions
			return
		}
	}

	// Have we gone too deep?
	if currentDepth == maxDepth {
		return
	}

	// If we are not done here, get satisfying reactions
	newPotentialReactions := rp.getSatisfyingReactions(substrateInputs, currentReactions)
	for _, potentialReaction := range newPotentialReactions {
		// Add a new reaction to the list
		currentReactions = append(currentReactions, potentialReaction)

		// Append to the substrateInputMap
		for product, _ := range rp.ReactionsToProductCompounds[potentialReaction] {
			substrateInputs = append(substrateInputs, product)
		}

		// Recurse
		wg.Add(1)
		go recurseReactionTree(rp, wg, validPathways, targetCompound, substrateInputs, currentReactions, currentDepth+1, maxDepth)
	}
}

func getPathways(validPathways chan []string, outputPathways chan [][]string) {
	var output [][]string
	for {
		pathway, more := <-validPathways
		if more {
			output = append(output, pathway)
		} else {
			outputPathways <- output
			close(outputPathways)
			return
		}
	}
}
