package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"flag"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"server/configs"
)

const projectID = "click-to-call-d2769"

func main() {
	var env = flag.String("env", "dev", "environment")
	flag.Parse()
	log.Printf("Main:%s", *env)
	config := configs.GetConfig(*env)
	ctx := context.Background()
	opt := config["firebaseOptions"].(option.ClientOption)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: projectID}, opt)
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

	http.HandleFunc("/logout", getLogoutHandler(ctx, firestoreClient, messagingClient))
	http.HandleFunc("/code", getCodeHandler(ctx, firestoreClient, messagingClient))
	http.HandleFunc("/call", getCallHandler(ctx, firestoreClient, messagingClient))
	http.HandleFunc("/healthy", HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}
