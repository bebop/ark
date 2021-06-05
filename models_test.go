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
	"testing"
)

var db *sqlx.DB
var minioClient *minio.Client

func TestMain(m *testing.M) {
	code := m.Run()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	options := &dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "latest",
		Cmd:        []string{"server", "/data"},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"9000/tcp": []dc.PortBinding{{HostPort: "9000"}},
		},
		Env: []string{"MINIO_ACCESS_KEY=MYACCESSKEY", "MINIO_SECRET_KEY=MYSECRETKEY"},
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

	// Shows buckets are set up

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	//// You can't defer this because os.Exit doesn't care for defer
	//if err := pool.Purge(resource); err != nil {
	//	log.Fatalf("Could not purge resource: %s", err)
	//}

	os.Exit(code)
}
