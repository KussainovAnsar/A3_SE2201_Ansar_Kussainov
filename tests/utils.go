package tests

import (
	"io"
	"log"
	"net/http"
)

func printResponseBody(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
}