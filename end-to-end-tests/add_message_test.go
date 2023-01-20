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
	createMessageUrl = basePath + "/messages/add"
)

type CreateMessageRequest struct {
	ChatId int    `json:"chat" binding:"required"`
	UserId int    `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type CreateMessageResponse struct {
	Id int `json:"messageId"`
}

func createMessageRequest(userId, chatId int, text string) (*http.Request, error) {
	message := CreateMessageRequest{
		ChatId: chatId,
		UserId: userId,
		Text:   text,
	}
	result, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", createMessageUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func TestAddMessageSuccess(t *testing.T) {
	// Arrange
	userId, err := createUser("user 5")
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat("chat 5", users)
	req, err := createMessageRequest(userId, chatId, "text")

	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var messageId CreateMessageResponse
	err = json.Unmarshal(body, &messageId)
	assert.GreaterOrEqual(t, messageId.Id, 1)
}

func TestAddMessageWithNotExistsUser(t *testing.T) {
	// Arrange
	users := make([]int, 0)
	chatId, err := createChat("chat 6", users)
	req, err := createMessageRequest(999, chatId, "text")

	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAddMessageWithNotExistsChat(t *testing.T) {
	// Arrange
	userId, err := createUser("user 7")
	assert.Nil(t, err)
	req, err := createMessageRequest(userId, 999, "text")

	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}



func TestAddMessageWithAuthorNotFromChat(t *testing.T) {
	// Arrange
	userId, err := createUser("user 8")
	assert.Nil(t, err)
	users := make([]int, 0)
	chatId, err := createChat("chat 9", users)
	req, err := createMessageRequest(userId, chatId, "text")

	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAddMessageWithEmptyBody(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("POST", createMessageUrl, bytes.NewBuffer([]byte("")))
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
