package end_to_end_tests

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	// Server
	host       = "webapp:9000"
	basePath   = "http://" + host
	healthPath = basePath + "/health"

	// Users
	createUserUrl = basePath + "/users/add"

	// Chats
	createChatUrl = basePath + "/chats/add"
	getChatsUrl   = basePath + "/chats/get"

	// Messages
	createMessageUrl = basePath + "/messages/add"
	getMessagesUrl   = basePath + "/messages/get"
)

const (
	NonExistentId = 999
)

func TestMain(m *testing.M) {
	err := healthCheck()
	if err != nil {
		log.Fatalf("Host %s is not available: %s", host, err)
	}

	log.Printf("Host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck() error {
	req, err := http.NewRequest("GET", healthPath, &bytes.Buffer{})
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("status code != 200 OK")
}

func getUniqueUserName() string {
	return "user-" + time.Now().Format(time.RFC3339Nano)
}

func getUniqueChatName() string {
	return "chat-" + time.Now().Format(time.RFC3339Nano)
}
