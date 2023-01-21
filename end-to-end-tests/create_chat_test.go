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

func TestCreateChatWithUserSuccess(t *testing.T) {
	// Arrange
	client := &http.Client{}

	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	req, err := createChatRequest(getUniqueChatName(), users)
	assert.Nil(t, err)

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

func TestCreateChatWithNotUniqueName(t *testing.T) {
	// Arrange
	client := &http.Client{}

	users := make([]int, 0)
	chatName := getUniqueChatName()
	req1, err := createChatRequest(chatName, users)
	assert.Nil(t, err)
	req2, err := createChatRequest(chatName, users)
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

func TestCreateChatWithTooLongName(t *testing.T) {
	// Arrange
	client := &http.Client{}

	users := make([]int, 0)
	req, err := createChatRequest(strings.Repeat("#", 100), users)
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateChatWithEmptyBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := http.NewRequest("POST", createChatUrl, bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateChatWithNonExistentUserId(t *testing.T) {
	// Arrange
	client := &http.Client{}

	users := make([]int, 1)
	users[0] = NonExistentId
	req, err := createChatRequest(getUniqueChatName(), users)
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateChatWithInvalidRequestBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	body := []byte(`"name": [1, 2, 3], "users": "123"`)
	req, err := http.NewRequest("POST", createChatUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
