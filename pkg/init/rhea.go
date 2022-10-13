package init

// TODO: uncomment and test this function.

// Rhea parses and inserts the rhea data into the database.
// func Rhea(ctx context.Context, db *surrealdb.DB, config config.Config) error {
// 	// parse Rhea file
// 	rheaBytes, err := rhea.ReadGzippedXml(config.RheaRDF)
// 	parsedRhea, err := rhea.Parse(rheaBytes)
// 	if err != nil {
// 		return err
// 	}

// 	// insert compounds
// 	for _, compound := range parsedRhea.Compounds {

// 		compoundID := "compound:" + strconv.Itoa(compound.ID) // TODO: remove if not needed
// 		// fmt.Println("parsed rhea", parsedRhea)
// 		_, err := db.Create(compoundID, compound)
// 		if err != nil {
// 			_, err := db.Update(compoundID, compound)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	// TODO: insert reaction participants and reaction sides
// 	for _, participant := range parsedRhea.ReactionParticipants {

// 		participantID := "participant:" + participant.Accession // TODO: remove if not needed
// 		_, err := db.Create(participantID, participant)
// 		if err != nil {
// 			_, err := db.Update(participantID, participant)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	// TODO: insert reactions
// 	for _, reaction := range parsedRhea.Reactions {
// 		reactionID := "reaction:" + strconv.Itoa(reaction.ID) // TODO: remove if not needed
// 		_, err := db.Create(reactionID, reaction)
// 		if err != nil {
// 			_, err := db.Update(reactionID, reaction)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
