package end_to_end_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createUserRequest(username string) (*http.Request, error) {
	user := CreateUserRequest{
		Username: username,
	}
	result, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func TestCreateUserSuccess(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := createUserRequest(getUniqueUserName())
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var userId CreateUserResponse
	err = json.Unmarshal(body, &userId)
	assert.GreaterOrEqual(t, userId.Id, 1)
}

func TestCreateUserWithNotUniqueName(t *testing.T) {
	// Arrange
	client := &http.Client{}

	username := getUniqueUserName()
	req1, err := createUserRequest(username)
	assert.Nil(t, err)
	req2, err := createUserRequest(username)
	assert.Nil(t, err)

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

func TestCreateUserWithTooLongName(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := createUserRequest(strings.Repeat("#", 100))
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithEmptyRequestBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUserWithInvalidRequestBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	body := []byte(`"username": [1, 2, 3]`)
	req, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
