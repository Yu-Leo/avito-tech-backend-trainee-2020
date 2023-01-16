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

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /users/add
func TestHTTPAddUserSuccess(t *testing.T) {
	body := `{
		"username": "name 9"
	}`
	var id int
	Test(t,
		Description("Successful user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusCreated),
		Store().Response().Body().JSON().JQ(".userId").In(&id),
	)
}

// HTTP POST: /users/add
func TestHTTPAddUserNotUniqueName(t *testing.T) {
	body := `{
		"username": "name 10"
	}`
	var id int
	Test(t,
		Description("Failed user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusCreated),
		Store().Response().Body().JSON().JQ(".userId").In(&id),
	)

	Test(t,
		Description("Failed user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

// HTTP POST: /users/add
func TestHTTPAddUserEmptyBody(t *testing.T) {
	body := `{}`
	Test(t,
		Description("Failed user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

// HTTP POST: /users/add
func TestHTTPAddUserInvalidRequest(t *testing.T) {
	body := `{
	"username": [1, 2, 3]
}`
	Test(t,
		Description("Failed user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}
