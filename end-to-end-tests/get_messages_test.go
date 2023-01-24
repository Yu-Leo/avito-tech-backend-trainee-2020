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
	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat(getUniqueChatName(), users)
	assert.Nil(t, err)
	messageId, err := createMessage(userId, chatId, "text")
	assert.Nil(t, err)

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

func TestGetChatMessagesCheckMessagesOrder(t *testing.T) {
	// Arrange
	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat(getUniqueChatName(), users)
	text1 := "text 1"
	message1Id, err := createMessage(userId, chatId, text1)
	assert.Nil(t, err)
	text2 := "text 1"
	message2Id, err := createMessage(userId, chatId, text2)
	assert.Nil(t, err)

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
	assert.Equal(t, len(chatMessages), 2)
	assert.Equal(t, chatMessages[0].Id, message1Id)
	assert.Equal(t, chatMessages[0].UserId, userId)
	assert.Equal(t, chatMessages[0].ChatId, chatId)
	assert.Equal(t, chatMessages[0].Text, text1)

	assert.Equal(t, chatMessages[1].Id, message2Id)
	assert.Equal(t, chatMessages[1].UserId, userId)
	assert.Equal(t, chatMessages[1].ChatId, chatId)
	assert.Equal(t, chatMessages[1].Text, text2)
}

func TestGetChatMessagesWithNonExistentChatId(t *testing.T) {
	// Arrange
	req, err := getMessagesRequest(NonExistentId)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetChatMessagesWithEmptyRequestBody(t *testing.T) {
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
