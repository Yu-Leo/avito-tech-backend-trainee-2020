package integration_tests

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "localhost:9000"
	healthPath = "http://" + host + "/health"
	attempts   = 20

	// HTTP REST
	basePath = "http://" + host
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for i := 0; i < attempts; i++ {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}
		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)
		time.Sleep(time.Second)
	}
	return err
}
