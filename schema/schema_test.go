package schema

import (
	//"context"
	//"fmt"

	"io/ioutil"
	"log"

	//"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/TimothyStiles/allbase/pkg/config"
	"github.com/jmoiron/sqlx"

	//"github.com/minio/minio-go/v7"
	//"github.com/minio/minio-go/v7/pkg/credentials"
	//"github.com/ory/dockertest/v3"
	//dc "github.com/ory/dockertest/v3/docker"

	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

//var minioClient *minio.Client

func TestCreateDatabase(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	tmpAllbaseConfig := config.DevDefault()
	tmpAllbaseConfig.AllbasePath = filepath.Join(tmpDataDir, "test.db")
	if err != nil {
		t.Errorf("Failed to create temporary data directory")
	}
	defer os.RemoveAll(tmpDataDir)

	err = CreateDatabase(tmpAllbaseConfig)

	if err != nil {
		log.Fatalf("Failed on error during database creation: %s", err)
	}
}

func TestMain(m *testing.M) {
	var err error
	//pool, err := dockertest.NewPool("")
	//if err != nil {
	//	log.Fatalf("Could not connect to docker: %s", err)
	//}

	//options := &dockertest.RunOptions{
	//	Repository:   "minio/minio",
	//	Tag:          "latest",
	//	Cmd:          []string{"server", "/data"},
	//	PortBindings: map[dc.Port][]dc.PortBinding{"9000/tcp": {{HostPort: "9000"}}},
	//	Env:          []string{"MINIO_ACCESS_KEY=MYACCESSKEY", "MINIO_SECRET_KEY=MYSECRETKEY"},
	//}

	//resource, err := pool.RunWithOptions(options)
	//if err != nil {
	//	log.Fatalf("Could not start resource: %s", err)
	//}

	//endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))
	//// or you could use the following, because we mapped the port 9000 to the port 9000 on the host
	//// endpoint := "localhost:9000"

	//// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	//// the minio client does not do service discovery for you (i.e. it does not check if connection can be established), so we have to use the health check
	//if err := pool.Retry(func() error {
	//	url := fmt.Sprintf("http://%s/minio/health/live", endpoint)
	//	resp, err := http.Get(url)
	//	if err != nil {
	//		return err
	//	}
	//	if resp.StatusCode != http.StatusOK {
	//		return fmt.Errorf("status code not OK")
	//	}
	//	resp.Body.Close()
	//	return nil
	//}); err != nil {
	//	log.Fatalf("Could not connect to docker: %s", err)
	//}

	//// now we can instantiate minio client
	//minioClient, err = minio.New(endpoint, &minio.Options{
	//	Creds:  credentials.NewStaticV4("MYACCESSKEY", "MYSECRETKEY", ""),
	//	Secure: false,
	//})
	//if err != nil {
	//	log.Println("Failed to create minio client:", err)
	//}

	//// now we can use the client, for example, to list the buckets
	//_, err = minioClient.ListBuckets(context.Background())
	//if err != nil {
	//	log.Fatalf("error while listing buckets: %v", err)
	//}

	// Begin SQLite
	db, err = sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open sqlite in memory: %s", err)
	}

	// Execute our schema in memory
	_, err = db.Exec(CreateSchema())
	if err != nil {
		log.Fatalf("Failed to execute schema: %s", err)
	}

	// Run the rest of our tests
	code := m.Run()

	//// You can't defer this because os.Exit doesn't care for defer
	//if err := pool.Purge(resource); err != nil {
	//	log.Fatalf("Could not purge resource: %s", err)
	//}

	os.Exit(code)
}

func TestChemblAttach(t *testing.T) {
	tmpDataDir, err := ioutil.TempDir("", "data-*")
	testConfig := config.DevDefault()

	if err != nil {
		t.Errorf("Failed to create temporary data directory")
	}
	defer os.RemoveAll(tmpDataDir)

	tmpChemblDBPath := filepath.Join(tmpDataDir, "chembl.db")

	// Read Chembl schema
	schemaStringBytes, err := ioutil.ReadFile(testConfig.ChemblSchema)
	if err != nil {
		t.Errorf("Failed to open chembl schema: %s", err)
	}

	// Begin SQLite
	chemblDB, err := sqlx.Open("sqlite3", tmpChemblDBPath)
	if err != nil {
		t.Errorf("Failed to open sqlite in %s: %s", tmpChemblDBPath, err)
	}

	// Execute our schema in memory
	_, err = chemblDB.Exec(string(schemaStringBytes))
	if err != nil {
		t.Errorf("Failed to execute schema: %s", err)
	}

	err = chemblAttach(db, tmpChemblDBPath)
	if err != nil {
		t.Errorf("Failed to attach chembl with error %s", err)
	}
}
