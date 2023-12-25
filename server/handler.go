package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getHandler(ctx context.Context, firestoreClient *firestore.Client) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
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
			writeCode(ctx, firestoreClient, "UGcFEEx4kUzS7Y90NogY", codePayload.Code)
		case http.MethodOptions:
			fmt.Fprintf(w, "Ok")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
