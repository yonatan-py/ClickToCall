package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

const projectID = "click-to-call-d2769"

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./click-to-call.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	firestoreClient, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Error initializing Firestore messagingClient: %v", err)
	}
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error getting Messaging messagingClient: %v\n", err)
	}

	http.HandleFunc("/code", getCodeHandler(ctx, firestoreClient, messagingClient))
	http.HandleFunc("/call", getCallHandler(ctx, messagingClient, firestoreClient))
	http.ListenAndServe(":8080", nil)
}
