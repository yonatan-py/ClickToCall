package main

import (
	"context"
	"net/http"
)

func main() {
	mainContext := context.Background()
	go startListeningForCalls(mainContext)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type CodePayload struct {
	Code string `json:"code"`
}
