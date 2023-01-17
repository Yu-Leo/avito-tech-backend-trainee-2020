package integration_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/stretchr/testify/assert"
)

type UserId struct {
	Id int `json:"userId"`
}

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
}

// HTTP POST: /users/add

func TestAddUserSuccess(t *testing.T) {
	// Arrange
	url := basePath + "/users/add"
	createUser := CreateUserDTO{
		Username: "name 1",
	}
	result, _ := json.Marshal(createUser)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var userId UserId
	err = json.Unmarshal(body, &userId)
	assert.GreaterOrEqual(t, userId.Id, 1)
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
