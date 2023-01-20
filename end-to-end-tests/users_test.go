package end_to_end_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	createUserUrl = basePath + "/users/add"
)

type CreateUserRequest struct {
	Id int `json:"userId"`
}

type CreateUserResponse struct {
	Username string `json:"username" binding:"required"`
}

// HTTP POST: /users/add

func newRequest(username string) (*http.Request, error) {
	user := CreateUserResponse{
		Username: username,
	}
	result, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func TestAddUserSuccess(t *testing.T) {
	// Arrange
	req, err := newRequest("user 1")
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var userId CreateUserRequest
	err = json.Unmarshal(body, &userId)
	assert.GreaterOrEqual(t, userId.Id, 1)
}

func TestAddUserNotUniqueName(t *testing.T) {
	// Arrange
	req1, err := newRequest("user 2")
	req2, err := newRequest("user 2")

	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res1, err := client.Do(req1)
	assert.Nil(t, err)
	defer res1.Body.Close()

	res2, err := client.Do(req2)
	assert.Nil(t, err)
	defer res2.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res1.StatusCode)
	assert.Equal(t, http.StatusBadRequest, res2.StatusCode)
}

func TestAddUserEmptyBody(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAddUserInvalidRequest(t *testing.T) {
	// Arrange
	body := []byte(`"username": [1, 2, 3]`)
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
