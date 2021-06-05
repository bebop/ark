package allbase

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/allyourbasepair/allbase/rhea"
	"github.com/koeng101/poly/parsers/uniprot"
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
	_, err = db.Exec(schema)
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

func TestRheaInsert(t *testing.T) {
	rhea, err := rhea.Read("rhea/data/rhea_mini.rdf.gz")
	if err != nil {
		log.Fatalf("Could not read rhea: %s", err)
	}

	err = RheaInsert(db, rhea)
	if err != nil {
		log.Fatalf("Could not insert rhea: %s", err)
	}
}

func TestUniprotInsert(t *testing.T) {
	var wg sync.WaitGroup
	uniprotSprot, errors, err := uniprot.ReadUniprot("data/uniprot_sprot_mini.xml.gz")
	if err != nil {
		log.Fatalf("Failed to read uniprot on error: %s", err)
	}
	wg.Add(1)
	go InsertUniprot(db, "sprot", uniprotSprot, errors, &wg)
	wg.Wait()

	for err := range errors {
		if err.Error() != "EOF" {
			log.Fatalf("Failed on error: %s", err)
		}
	}
}
