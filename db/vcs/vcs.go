package vcs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type ActionType string // ActionType is the type of action that was performed on a row in a table

const (
	Initialize ActionType = "initialize"
	Insert     ActionType = "insert"
	Modify     ActionType = "modify"
	Delete     ActionType = "delete"
)

// Row is a struct that contains the information about a row in a table that we need to pass through VCS
type Row struct {
	PrimaryKeyColumn string
	PrimaryKeyValue  interface{}
	Fields           map[string]interface{}
	TableName        string
}

// Change is a struct that contains the information about a change made to a row in a table
type Change struct {
	ID                  uint64      `db:"id"`
	Action              string      `db:"action"`
	RowPrimaryKeyColumn string      `db:"row_primary_key_column"`
	RowPrimaryKeyValue  interface{} `db:"row_primary_key_value"`
	Field               string      `db:"field"`
	OldData             interface{} `db:"old_data"`
	NewData             interface{} `db:"new_data"`
	ForkID              uint64      `db:"fork_id"`
	TableName           string      `db:"table_name"`
	CreatedAt           time.Time   `db:"created_at"`
}

type Fork struct {
	ID             int           `db:"id"`
	ParentID       sql.NullInt64 `db:"parent_id"`
	Name           string        `db:"name"`
	CreatedAt      time.Time     `db:"created_at"`
	ChangeOriginID uint64        `db:"change_origin_id"`
}

type VCSState struct {
	CurrentForkID   uint64 `json:"current_fork_id"`
	CurrentChangeID uint64 `json:"current_change_id"`
}

type Changes []*Change

type VCS struct {
	SynbioDB *sqlx.DB
	VCSDB    *sqlx.DB
	VCSState *VCSState
}

// Create a connection to VCS Sqlite3 database
func ConnectVCSDB() (*sqlx.DB, error) {
	// Get the path to the VCS Sqlite3 database from the environment
	dbPath, ok := os.LookupEnv("VCS_DB_PATH")
	if !ok {
		dbPath = "./vcs.db"
	}

	// Open a connection to the VCS Sqlite3 database
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create a connection to the SynBio Sqlite3 database
func ConnectSynBioDB() (*sqlx.DB, error) {
	// Get the path to the SynBio Sqlite3 database from the environment
	dbPath, ok := os.LookupEnv("SYNBIO_DB_PATH")
	if !ok {
		dbPath = "../data/dev/retsynth/minimal.db"
	}

	// Open a connection to the SynBio Sqlite3 database
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Find the existing row in the VCS database
func findExistingRow(row *Row) *Row {
	// Connect to the SynBio Database
	db, err := ConnectSynBioDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	// Query the Row using the given table
	rawrow := db.QueryRow("SELECT * FROM ? WHERE ? = ?", row.TableName, row.PrimaryKeyColumn, row.PrimaryKeyValue)

	// Check if the row exists
	if rawrow == nil {
		return nil
	}

	// Map the row to a new Row struct
	var oldRow Row
	oldRow.PrimaryKeyColumn = row.PrimaryKeyColumn
	oldRow.PrimaryKeyValue = row.PrimaryKeyValue
	oldRow.Fields = make(map[string]interface{})
	oldRow.TableName = row.TableName
	for field, value := range row.Fields {
		oldRow.Fields[field] = value
	}

	return &oldRow
}

// Generate a list of changes made to the data given the rows that we are making changes to the data (CREATE, UPDATE)
func TrackChanges(data []*Row, deletedData []*Row) (Changes, error) {
	// Initialize a list of changes to track the changes made to the data
	changes := Changes{}

	// Iterate through the rows in the current state of the data
	for _, currentRow := range data {
		// Check if the current row exists in the previous state of the data
		previousRow := findExistingRow(currentRow)

		// If the current row does not exist in the previous state, it has been added
		if previousRow == nil {
			for field, value := range currentRow.Fields {
				changes = append(
					changes,
					&Change{
						Action:              string(Insert),
						RowPrimaryKeyColumn: currentRow.PrimaryKeyColumn,
						RowPrimaryKeyValue:  currentRow.PrimaryKeyValue,
						TableName:           previousRow.TableName,
						Field:               field,
						OldData:             previousRow.Fields[field],
						NewData:             value,
					})
			}
		} else {
			// If the current row does exist in the previous state, check if it has been modified
			// Iterate through the fields in the current row
			for field, value := range currentRow.Fields {
				// If the field has a different value in the previous state, it has been modified
				if previousRow.Fields[field] != value {
					changes = append(
						changes,
						&Change{
							Action:              string(Modify),
							RowPrimaryKeyColumn: currentRow.PrimaryKeyColumn,
							RowPrimaryKeyValue:  currentRow.PrimaryKeyValue,
							TableName:           previousRow.TableName,
							Field:               field, OldData: previousRow.Fields[field],
							NewData: value,
						})
				}
			}
		}
	}

	// Iterate through the deleted rows and generate a list of changes
	for _, deletedRow := range deletedData {
		// Check if the deleted row exists in the previous state of the data
		previousRow := findExistingRow(deletedRow)
		// If the deleted row does not exist, its an error
		if previousRow == nil {
			errormessage := fmt.Sprintf("deleted row in table: %s with primary key (%s) and field %s does not exist", deletedRow.TableName, deletedRow.PrimaryKeyValue, deletedRow.PrimaryKeyColumn)
			return nil, errors.New(errormessage)
		}
		// If the deleted row does exist, generate a list of changes
		changes = append(
			changes,
			&Change{
				Action:              string(Delete),
				RowPrimaryKeyColumn: previousRow.PrimaryKeyColumn,
				RowPrimaryKeyValue:  previousRow.PrimaryKeyValue,
				TableName:           previousRow.TableName,
				NewData:             nil,
			})

		// Iterate through all the fields and create a modify entry for each field
		for field, value := range previousRow.Fields {
			changes = append(
				changes,
				&Change{
					Action:              string(Modify),
					RowPrimaryKeyColumn: previousRow.PrimaryKeyColumn,
					RowPrimaryKeyValue:  previousRow.PrimaryKeyValue,
					TableName:           previousRow.TableName,
					Field:               field,
					OldData:             value,
					NewData:             nil,
				})
		}
	}

	return changes, nil
}

// Get the changes made to the data for a given fork
func GetChanges(forkID int) ([]Change, error) {
	// Connect to the VCS Database
	db, err := ConnectVCSDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Construct sqlx query to get all changes for a given fork
	query := "SELECT * FROM changes WHERE fork_id = ?"
	var changes []Change
	err = db.Select(&changes, query, forkID)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

// Rollback the changes made to the data to a given index in the list of changes
func RollbackTo(changeID int, forkID int) error {
	// Get the list of changes made to the data in the VCS database for the given fork
	changes, err := GetChanges(forkID)
	if err != nil {
		return err
	}

	if changeID < 0 || changeID >= len(changes) {
		return errors.New("invalid index")
	}

	// Connect to the SynBio Database
	db, err := ConnectSynBioDB()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for i := len(changes) - 1; i > changeID; i-- {
		change := changes[i]
		if change.Action == string(Insert) {
			if _, err := tx.Exec("DELETE FROM ? WHERE ? = ?", change.TableName, change.RowPrimaryKeyColumn, change.RowPrimaryKeyValue); err != nil {
				tx.Rollback()
				return err
			}
		} else if change.Action == string(Modify) {
			if _, err := tx.Exec("UPDATE ? SET ? = ? WHERE ? = ?", change.TableName, change.Field, change.OldData, change.RowPrimaryKeyColumn, change.RowPrimaryKeyValue); err != nil {
				tx.Rollback()
				return err
			}
		} else if change.Action == string(Delete) {
			_, err := tx.Exec("INSERT INTO ? (?) VALUES (?)", change.TableName, change.RowPrimaryKeyColumn, change.RowPrimaryKeyValue)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

// Get the current state of the VCS database
func ReadVCSState() (VCSState, error) {
	// Read the vcs state from the json file
	data, err := os.ReadFile("vcs.json")
	if err != nil {
		return VCSState{}, err
	}

	// Unmarshal the json data into a VCSState struct
	var vcsState VCSState
	err = json.Unmarshal(data, &vcsState)
	if err != nil {
		return VCSState{}, err
	}

	return vcsState, nil
}

// Write the VCS state to the vcs.json file
func WriteVCSState(vcsState VCSState) error {
	// Marshal the VCSState struct into json
	data, err := json.Marshal(vcsState)
	if err != nil {
		return err
	}

	// Write the json data to the vcs.json file
	err = os.WriteFile("vcs.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Create a fork and return the ID of the fork
func CreateFork(forkname string) (uint64, error) {

	//Read the vcs state of the database
	data, err := ReadVCSState()
	if err != nil {
		return 0, err
	}

	// Connect to the VCS Database
	db, err := ConnectVCSDB()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("INSERT INTO forks (name, parent_id, created_at, change_origin_id) VALUES (?, ?, ?, ?)", forkname, data.CurrentForkID, time.Now(), data.CurrentChangeID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
