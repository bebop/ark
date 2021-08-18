package models

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/TimothyStiles/poly"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"

	"github.com/TimothyStiles/poly/parsers/uniprot"
	"github.com/allyourbasepair/allbase/rhea"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB
var minioClient *minio.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	options := &dockertest.RunOptions{
		Repository:   "minio/minio",
		Tag:          "latest",
		Cmd:          []string{"server", "/data"},
		PortBindings: map[dc.Port][]dc.PortBinding{"9000/tcp": {{HostPort: "9000"}}},
		Env:          []string{"MINIO_ACCESS_KEY=MYACCESSKEY", "MINIO_SECRET_KEY=MYSECRETKEY"},
	}

	resource, err := pool.RunWithOptions(options)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))
	// or you could use the following, because we mapped the port 9000 to the port 9000 on the host
	// endpoint := "localhost:9000"

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	// the minio client does not do service discovery for you (i.e. it does not check if connection can be established), so we have to use the health check
	if err := pool.Retry(func() error {
		url := fmt.Sprintf("http://%s/minio/health/live", endpoint)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code not OK")
		}
		resp.Body.Close()
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// now we can instantiate minio client
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("MYACCESSKEY", "MYSECRETKEY", ""),
		Secure: false,
	})
	if err != nil {
		log.Println("Failed to create minio client:", err)
	}

	// now we can use the client, for example, to list the buckets
	_, err = minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("error while listing buckets: %v", err)
	}

	// Begin SQLite
	db, err = sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open sqlite in memory: %s", err)
	}

	// Execute our schema in memory
	_, err = db.Exec(Schema)
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	// Run the rest of our tests
	code := m.Run()

	//// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestUniprotInsert(t *testing.T) {
	// First, test Rhea insert. We need both to test uniprot2rhea
	rhea, err := rhea.Read("../rhea/data/rhea_mini.rdf.gz")
	if err != nil {
		log.Fatalf("Could not read rhea: %s", err)
	}

	err = RheaInsert(db, rhea)
	if err != nil {
		log.Fatalf("Could not insert rhea: %s", err)
	}

	// Then uniprot
	var wg sync.WaitGroup
	uniprotSprot, errors, err := uniprot.Read("../data/uniprot_sprot_mini.xml.gz")
	if err != nil {
		log.Fatalf("Failed to read uniprot on error: %s", err)
	}
	wg.Add(1)
	go UniprotInsert(db, "sprot", uniprotSprot, errors, &wg)
	wg.Wait()

	for err := range errors {
		if err.Error() != "EOF" {
			log.Fatalf("Failed on error during uniprot parsing or insertion: %s", err)
		}
	}

	// Finally, UniprotToRhea
	err = RheaTsvInsert(db, "../data/rhea2uniprot_test.tsv.gz", true)
	if err != nil {
		log.Fatalf("Failed to insert RheaTsvInsert on: %s", err)
	}
}

func TestGenbankInsert(t *testing.T) {
	sequences := poly.ReadGbkFlatGz("../data/flatGbk_test.seq.gz")
	err := GenbankInsert(db, sequences)
	if err != nil {
		log.Fatalf("Failed on error during genbank insertion: %s", err)
	}
}

func TestChemblAttach(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	if err != nil {
		t.Errorf("Failed to create temporary data directory")
	}
	defer os.RemoveAll(tmpDataDir)

	tmpChemblFilePath := filepath.Join(tmpDataDir, "chembl.db")

	// Read Chembl schema
	schemaStringBytes, err := ioutil.ReadFile("../data/chembl_schema.sql")
	if err != nil {
		t.Errorf("Failed to open chembl schema: %s", err)
	}

	// Begin SQLite
	chemblDB, err := sqlx.Open("sqlite3", tmpChemblFilePath)
	if err != nil {
		t.Errorf("Failed to open sqlite in %s: %s", tmpChemblFilePath, err)
	}

	// Execute our schema in memory
	_, err = chemblDB.Exec(string(schemaStringBytes))
	if err != nil {
		t.Errorf("Failed to execute schema: %s", err)
	}

	err = ChemblAttach(db, tmpChemblFilePath)
	if err != nil {
		t.Errorf("Failed to attach chembl with error %s", err)
	}
}
