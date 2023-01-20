package end_to_end_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	getMessagesUrl = basePath + "/messages/get"
)

type GetChatMessagesRequest struct {
	ChatId int `json:"chat" binding:"required"`
}

type GetChatMessagesResponse struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func getMessagesRequest(chatId int) (*http.Request, error) {
	message := GetChatMessagesRequest{
		ChatId: chatId,
	}
	result, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", getMessagesUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func createMessage(userId, chatId int, text string) (int, error) {
	// Arrange
	req, err := createMessageRequest(userId, chatId, text)
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
	var messageId CreateMessageResponse
	err = json.Unmarshal(body, &messageId)
	return messageId.Id, nil
}

func TestGetChatMessagesSuccess(t *testing.T) {
	// Arrange
	userId, err := createUser("user 11")
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat("chat 8", users)
	messageId, err := createMessage(userId, chatId, "text")

	req, err := getMessagesRequest(chatId)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var chatMessages []GetChatMessagesResponse
	err = json.Unmarshal(body, &chatMessages)
	assert.Equal(t, len(chatMessages), 1)
	assert.Equal(t, chatMessages[0].Id, messageId)
	assert.Equal(t, chatMessages[0].UserId, userId)
	assert.Equal(t, chatMessages[0].ChatId, chatId)
	assert.Equal(t, chatMessages[0].Text, "text")
}

func TestGetChatMessagesNotExistsChat(t *testing.T) {
	// Arrange
	req, err := getMessagesRequest(999)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetChatMessagesWithEmptyBody(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("POST", getMessagesUrl, bytes.NewBuffer([]byte("{}")))
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
