package integration_tests

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"testing"
)

const (
	// Attempts connection
	host       = "webapp:9000"
	healthPath = "http://" + host + "/health"
	attempts   = 20

	// HTTP REST
	basePath = "http://" + host
)

func TestMain(m *testing.M) {
	err := healthCheck()
	if err != nil {
		log.Fatalf("Host %s is not available: %s", host, err)
	}

	log.Printf("Host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck() error {
	req, err := http.NewRequest("GET", healthPath, &bytes.Buffer{})
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("status code != 200 OK")
}
