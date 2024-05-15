package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime string   `json:"runtime"`
	Genres  []string `json:"genres"`
}
func TestCreateMovie(t *testing.T) {
	try(func() {
		moviePayload := Movie{
			Title:   "The Dark Knight",
			Year:    2008,
			Runtime: "152 mins",
			Genres:  []string{"action", "crime", "drama"},
		}

		payloadBytes, err := json.Marshal(moviePayload)
		if err != nil {
			throw("Error marshaling movie payload", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
		if err != nil {
			throw("Error creating HTTP request", err)
		}

		bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU" // authentication_token for my user
		req.Header.Set("Authorization", "Bearer "+bearerToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			throw("Error sending HTTP request", err)
		}
		defer res.Body.Close()

		printResponseBody(res)
	})
}

func TestCreateMovie_InvalidTitle(t *testing.T) {
	try(func() {
		moviePayload := Movie{
			Title:   "",
			Year:    2015,
			Runtime: "151 mins",
			Genres:  []string{"crime"},
		}

		payloadBytes, err := json.Marshal(moviePayload)
		if err != nil {
			throw("Error marshaling movie payload", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
		if err != nil {
			throw("Error creating HTTP request", err)
		}

		bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
		req.Header.Set("Authorization", "Bearer "+bearerToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			throw("Error sending HTTP request", err)
		}
		defer res.Body.Close()

		printResponseBody(res)
	})
}

func TestCreateMovie_InvalidYear(t *testing.T) {
	try(func() {
		moviePayload := Movie{
			Title:   "The Dark Knight",
			Year:    2050,
			Runtime: "152 mins",
			Genres:  []string{"action", "crime", "drama"},
		}

		payloadBytes, err := json.Marshal(moviePayload)
		if err != nil {
			throw("Error marshaling movie payload", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
		if err != nil {
			throw("Error creating HTTP request", err)
		}

		bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
		req.Header.Set("Authorization", "Bearer "+bearerToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			throw("Error sending HTTP request", err)
		}
		defer res.Body.Close()

		printResponseBody(res)
	})
}

func TestDeleteMovie(t *testing.T) {
	try(func() {
		req, err := http.NewRequest("DELETE", "http://localhost:4000/v1/movies/1", nil)
		if err != nil {
			throw("Error creating HTTP request", err)
		}

		bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
		req.Header.Set("Authorization", "Bearer "+bearerToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			throw("Error sending HTTP request", err)
		}
		defer res.Body.Close()

		printResponseBody(res)
	})
}

func try(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Error occurred: %v", err)
		}
	}()
	fn()
}

func throw(msg string, err error) {
	panic(msg + ": " + err.Error())
}