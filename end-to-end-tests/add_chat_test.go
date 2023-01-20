package end_to_end_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	createChatUrl = basePath + "/chats/add"
)

type CreateChatRequest struct {
	Name  string `json:"name" binding:"required"`
	Users []int  `json:"users" binding:"required"`
}

type CreateChatResponse struct {
	Id int `json:"chatId"`
}

func createChatRequest(name string, users []int) (*http.Request, error) {
	chat := CreateChatRequest{
		Name:  name,
		Users: users,
	}
	result, _ := json.Marshal(chat)
	req, err := http.NewRequest("POST", createChatUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func createUser(username string) (int, error) {
	// Arrange
	req, err := createUserRequest(username)
	if err != nil {
		return 0, err
	}
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	// Assert
	if res.StatusCode != http.StatusCreated {
		return 0, errors.New("invalid status code")
	}
	body, _ := io.ReadAll(res.Body)
	var userId CreateUserResponse
	err = json.Unmarshal(body, &userId)
	return userId.Id, nil
}

func TestAddChatWithUserSuccess(t *testing.T) {
	// Arrange
	userId, err := createUser("user 3")
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	req, err := createChatRequest("chat 1", users)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var chatId CreateChatResponse
	err = json.Unmarshal(body, &chatId)
	assert.GreaterOrEqual(t, chatId.Id, 1)
}

func TestAddChatWithNotUniqueName(t *testing.T) {
	// Arrange
	users := make([]int, 0)
	req1, err := createChatRequest("chat 2", users)
	req2, err := createChatRequest("chat 2", users)
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

func TestAddChatWithEmptyBody(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("POST", createChatUrl, bytes.NewBuffer([]byte("{}")))
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

func TestAddChatWithNotExistsUser(t *testing.T) {
	// Arrange
	var users []int
	users = append(users, 9999)
	req, err := createChatRequest("chat 3", users)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAddChatWithInvalidRequest(t *testing.T) {
	// Arrange
	body := []byte(`"name": [1, 2, 3], "users": "123"`)
	req, err := http.NewRequest("POST", createChatUrl, bytes.NewBuffer(body))
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
