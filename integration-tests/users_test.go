package integration_tests

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

// HTTP POST: /users/add

func TestAddUserSuccess(t *testing.T) {
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

func TestAddUserNotUniqueName(t *testing.T) {
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

func TestAddUserEmptyBody(t *testing.T) {
	body := `{}`
	Test(t,
		Description("Failed user addition"),
		Post(basePath+"/users/add"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

func TestAddUserInvalidRequest(t *testing.T) {
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
