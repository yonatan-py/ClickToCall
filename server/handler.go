package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "chrome-extension://hnedcaamblfpdpkblfocnjchfbianddf")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	println("handling")

	switch request.Method {
	case http.MethodPost:
		var codePayload CodePayload
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&codePayload)
		if err != nil {
			panic(err)
		}
		log.Println(codePayload)
		fmt.Fprintf(w, "Hi there, I love!")
	case http.MethodOptions:
		fmt.Fprintf(w, "Ok")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
