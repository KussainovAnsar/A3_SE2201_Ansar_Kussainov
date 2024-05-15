package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Ansar",
		"Email":    "kitsukidame@gmail.com",
		"Password": "Password123456",
	})

	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Ansar",
		"Email":    "hsjdhjdhsj",
		"Password": "Password123456"})

	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestGetAuthenticationToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "kitsukidame@gmail.com",
		"Password": "Password123456",
	})
	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestGetAuthenticationToken_InvalidBody(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "someerror@error.com",
		"Password": "Password123456",
	})

	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestActivateAccount(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"token": "ILXDEAQ3DLCEUQ2X72NJTXCC2M",
	})

	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", responseBody)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestActivateAccount_InvalidToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"token": "56VZQBTITZGF7R3L5QPGBRJH6M", // Invalid Activation Token
	})

	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", responseBody)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}
