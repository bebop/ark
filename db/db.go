package db

// Database is an interface that all database backends should implement.
type Database interface {

	// Connect should open a connection to the database.
	Connect() error

	// Init should create the database if it doesn't exist.
	Init() error

	// Config should return the database configuration.
	Config() (Config, error)

	// SetConfig should set the database configuration.
	SetConfig(Config) error

	// // Get returns the value for the given key.
	// Get(key string) (string, error)

	// // Set sets the value for the given key.
	// Set(key, value string) error

	// Delete removes the value for the given key.
	Delete(key string) error

	// Query returns a list of key/value pairs for the given query.
	Query(query string) ([]map[string]string, error)

	// Close closes the database.
	Close() error
}

// Config is the configuration for a database.
type Config struct {
	// Driver is the database driver.
	Driver string

	// Host is the hostname of the database.
	Host string

	// Port is the port of the database.
	Port int

	// Username is the username for the database.
	Username string

	// Password is the password for the database.
	Password string

	// Name is the name of the database.
	Name string

	// SSL is whether or not to use SSL.
	SSL bool
}

type query struct {
	language string
	query    string
}

// Pathway stores metabolic pathway and reaction information.
type Pathway struct {
	// ID is the pathway ID.
	ID string

	// Name is the pathway name.
	Name string

	// Description is the pathway description.
	Description string

	// Reactions is a list of reactions in the pathway.
	Reactions []Reaction
}

// DNA stores DNA sequence information.
type DNA struct {
}

// Protein stores protein sequence information.
type Protein struct {
}

// Reaction stores information about a chemical reaction.
type Reaction struct {
	// ID is the reaction ID.
	ID         string
	Upstream   *Reagent
	Downstream *Reagent
}

// Reagent stores information about a chemical reagent.
type Reagent struct {
	// ID is the reagent ID.
	ID string
}
