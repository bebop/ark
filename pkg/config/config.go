package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/TimothyStiles/allbase/pkg/env"
)

// Read reads the config file and returns the config.
func Read(fileName string) (Config, error) {
	configPath := filepath.Join(env.RootPath(), fileName)
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		return Config{}, err
	}

	var config Config
	json.NewDecoder(configFile).Decode(&config)
	return config, nil
}

// Write writes the config to the given file.
func Write(fileName string, config Config) error {
	configPath := filepath.Join(env.RootPath(), fileName)
	configFile, _ := os.Create(configPath)
	defer configFile.Close()

	json.NewEncoder(configFile).Encode(config)
	return nil
}

// Config is the configuration for the allbase application.
type Config struct {
	// Production or development.
	IsProd bool `json:"prod"`

	// RootPath is the root directory of the project.
	RootPath string `json:"root_path"`

	// DataPath is the data directory of the project.
	DataPath string `json:"data_path"`

	// AllbasePath is the path to the allbase sqlite database.
	AllbaseURL string `json:"allbase_URL"`

	// AdminUser is the admin user for the allbase database.
	AdminUser string `json:"admin_user"`

	// AdminPassword is the admin password for the allbase database.
	AdminPassword string `json:"admin_password"`

	// DBName is the name of the allbase database.
	DBName string `json:"db_name"`

	// RheaRDF is the path to the Rhea RDF file.
	RheaRDF string `json:"rhea_rdf"`

	// RheaToUniprotSprot is the path to the Rhea to Uniprot Sprot mapping file.
	RheaToUniprotSprot string `json:"rhea_to_uniprot_sprot"`

	// RheaToUniprotTrembl is the path to the Rhea to Uniprot Trembl mapping file.
	RheaToUniprotTrembl string `json:"rhea_to_uniprot_trembl"`

	// ChemblSchema is the path to the Chembl schema file.
	ChemblSchema string `json:"chembl_schema"`

	// ChemblSQLite is the path to the CHEMBL sqlite file.
	ChemblSQLite string `json:"chembl_sqlite"`

	// UniprotSprotXML is the path to the Uniprot Sprot XML file.
	UniprotSprotXML string `json:"uniprot_sprot_xml"`

	// UniprotTremblXML is the path to the Uniprot Trembl XML file.
	UniprotTremblXML string `json:"uniprot_trembl_xml"`

	// Genbank is the path to the Genbank directory.
	Genbank string `json:"genbank"`
}

// DevDefault returns the default configuration for development.
func DevDefault() Config {

	devPath := filepath.Join(env.RootPath(), "data", "dev")
	chemblSchemaPath := filepath.Join(env.RootPath(), "data", "chembl_schema.sql")
	return Config{
		IsProd:              false,
		RootPath:            env.RootPath(),
		DataPath:            devPath,
		AllbaseURL:          "ws://localhost:8000/rpc",
		AdminUser:           "root", // TODO: Change this.
		AdminPassword:       "root",
		DBName:              "allbase",
		RheaRDF:             filepath.Join(devPath, "rhea_mini.rdf.gz"),
		RheaToUniprotSprot:  filepath.Join(devPath, "rhea_to_uniprot_sprot.tsv.gz"),
		RheaToUniprotTrembl: filepath.Join(devPath, "rhea_to_uniprot_trembl.tsv.gz"),
		ChemblSchema:        chemblSchemaPath,
		ChemblSQLite:        filepath.Join(devPath, "chembl.sqlite"),
		UniprotSprotXML:     filepath.Join(devPath, "uniprot_sprot_test.xml.gz"),
		UniprotTremblXML:    filepath.Join(devPath, "uniprot_sprot_test.xml.gz"),
		Genbank:             filepath.Join(devPath, "genbank"),
	}
}

// ProdDefault returns the default configuration for production.
func ProdDefault() Config {
	prodPath := filepath.Join(env.RootPath(), "data", "prod")
	chemblSchemaPath := filepath.Join(env.RootPath(), "data", "chembl_schema.sql")

	return Config{
		IsProd:              true,
		RootPath:            env.RootPath(),
		DataPath:            prodPath,
		AllbaseURL:          "", // Sysadmin to fill in.
		AdminUser:           "", // Sysadmin to fill in.
		AdminPassword:       "", // Sysadmin to fill in.
		DBName:              "allbase",
		RheaRDF:             filepath.Join(prodPath, "rhea.rdf.gz"),
		RheaToUniprotSprot:  filepath.Join(prodPath, "rhea_to_uniprot_sprot.tsv"),
		RheaToUniprotTrembl: filepath.Join(prodPath, "rhea_to_uniprot_trembl.tsv"),
		ChemblSchema:        chemblSchemaPath,
		ChemblSQLite:        filepath.Join(prodPath, "chembl.sqlite"),
		UniprotSprotXML:     filepath.Join(prodPath, "uniprot_sprot.xml.gz"),
		UniprotTremblXML:    filepath.Join(prodPath, "uniprot_trembl.xml.gz"),
		Genbank:             filepath.Join(prodPath, "genbank"),
	}
}

// TestDefault returns the default configuration for testing.
func TestDefault() Config {
	return DevDefault()
}
