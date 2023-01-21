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

func getChatsRequest(user int) (*http.Request, error) {
	chat := GetUserChatsRequest{
		User: user,
	}
	result, _ := json.Marshal(chat)
	req, err := http.NewRequest("POST", getChatsUrl, bytes.NewBuffer(result))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func createUser(username string) (int, error) {
	// Arrange
	client := &http.Client{}

	req, err := createUserRequest(username)
	if err != nil {
		return 0, err
	}

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

func createChat(name string, users []int) (int, error) {
	// Arrange
	client := &http.Client{}

	req, err := createChatRequest(name, users)
	if err != nil {
		return 0, err
	}

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
	client := &http.Client{}

	userId, err := createUser(getUniqueUserName())
	assert.Nil(t, err)
	users := make([]int, 1)
	users[0] = userId
	chatName := getUniqueChatName()
	chatId, err := createChat(chatName, users)
	assert.Nil(t, err)
	req, err := getChatsRequest(userId)
	assert.Nil(t, err)

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
	assert.Equal(t, userChats[0].Name, chatName)
}

func TestGetUserChatsWithNonExistentUserId(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := getChatsRequest(NonExistentId)
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetUserChatsWithEmptyRequestBody(t *testing.T) {
	// Arrange
	client := &http.Client{}

	req, err := http.NewRequest("POST", getChatsUrl, bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.Nil(t, err)

	// Act
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
