package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
)

func listenDocument(ctx context.Context, w io.Writer, projectID, collection string, callback func(map[string]interface{})) error {

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("firestore.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Collection(collection).Doc("UGcFEEx4kUzS7Y90NogY").Snapshots(ctx)
	for {
		snap, err := it.Next()
		if status.Code(err) == codes.DeadlineExceeded {
			return nil
			if err != nil {
			}

			return fmt.Errorf("Snapshots.Next: %w", err)
		}
		if !snap.Exists() {
			fmt.Fprintf(w, "Document no longer exists\n")
			return nil
		}

		callback(snap.Data())
	}
}

func sendMessage(ctx context.Context, client messaging.Client, registrationToken string, number string) {
	message := &messaging.Message{
		Data: map[string]string{
			"number": number,
		},
		Token: registrationToken,
		Notification: &messaging.Notification{
			Title: "click to call",
			Body:  number,
		},
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", response)
}

func main() {

	opt := option.WithCredentialsFile("./click-to-call-d2769-2021c5c952d8.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)

	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	registrationToken := "cbWMXrTYRg-VQQSoPpRq7T:APA91bGYxuLIdfvTTik_0FVhJ3yG5djtRSUZ6sHgXRBrVd0gGbDskydRDwamNTiUztpk9oc25oXwm3-AviioahOlwAJTn6cxRzxdcFPG3O37Rus2p6RiI6nSiYkVqb4kaYY4FC56cdZM"

	projectID := "click-to-call-d2769"

	notifyChange := func(data map[string]interface{}) {
		number := data["call"].(string)
		println(number)
		sendMessage(ctx, *client, registrationToken, number)
	}
	listenDocument(ctx, os.Stdout, projectID, "users", notifyChange)
}
