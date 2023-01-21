package end_to_end_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCreateMessageSuccess(t *testing.T) {
	// Arrange
	client := &http.Client{}

	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat(getUniqueChatName(), users)
	assert.Nil(t, err)
	req, err := createMessageRequest(userId, chatId, "text")
	assert.Nil(t, err)

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

func TestCreateMessageWithNonExistentUserId(t *testing.T) {
	// Arrange
	client := &http.Client{}

	users := make([]int, 0)
	chatId, err := createChat(getUniqueChatName(), users)
	assert.Nil(t, err)

	req, err := createMessageRequest(NonExistentId, chatId, "text")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateMessageWithNonExistentChatId(t *testing.T) {
	// Arrange
	client := &http.Client{}

	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	req, err := createMessageRequest(userId, NonExistentId, "text")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateMessageWithAuthorNotFromChat(t *testing.T) {
	// Arrange
	client := &http.Client{}

	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 0)
	chatId, err := createChat(getUniqueChatName(), users)
	assert.Nil(t, err)
	req, err := createMessageRequest(userId, chatId, "text")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateMessageWithEmptyRequestBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := http.NewRequest("POST", createMessageUrl, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
