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
	getChatsUrl = basePath + "/chats/get"
)

type GetUserChatsRequest struct {
	User int `json:"user" binding:"required"`
}

type GetUserChatsResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []int     `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

func getChatsRequest(user int) (*http.Request, error) {
	chat := GetUserChatsRequest{
		User: user,
	}
	result, _ := json.Marshal(chat)
	req, err := http.NewRequest("POST", getChatsUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func createChat(name string, users []int) (int, error) {
	// Arrange
	req, err := createChatRequest(name, users)
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
	var chatId CreateChatResponse
	err = json.Unmarshal(body, &chatId)
	return chatId.Id, nil
}

func TestGetUserChatsSuccess(t *testing.T) {
	// Arrange
	userId, err := createUser("user 4")
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatId, err := createChat("chat 4", users)
	req, err := getChatsRequest(userId)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	var userChats []GetUserChatsResponse
	err = json.Unmarshal(body, &userChats)
	assert.Equal(t, len(userChats), 1)
	assert.Equal(t, userChats[0].Id, chatId)
	assert.Equal(t, userChats[0].Name, "chat 4")
}

func TestGetUserChatsWithNotExistsUser(t *testing.T) {
	// Arrange
	req, err := getChatsRequest(999)
	assert.Nil(t, err)
	client := &http.Client{}

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetUserChatsWithEmptyBody(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("POST", getChatsUrl, bytes.NewBuffer([]byte("{}")))
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
