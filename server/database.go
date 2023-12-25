package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

const registrationToken = "cbWMXrTYRg-VQQSoPpRq7T:APA91bGYxuLIdfvTTik_0FVhJ3yG5djtRSUZ6sHgXRBrVd0gGbDskydRDwamNTiUztpk9oc25oXwm3-AviioahOlwAJTn6cxRzxdcFPG3O37Rus2p6RiI6nSiYkVqb4kaYY4FC56cdZM"

func startListeningForCalls(ctx context.Context, app *firebase.App, projectID string) {

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	notifyChange := func(data map[string]interface{}) {
		number := data["call"].(string)
		println(number)
		sendMessage(ctx, *client, number)
	}
	err = listenDocument(ctx, projectID, notifyChange)
	if err != nil {
		return
	}
}

func sendMessage(ctx context.Context, client messaging.Client, number string) {
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

func listenDocument(ctx context.Context, projectID string, callback func(map[string]interface{})) error {

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("firestore.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Collection("users").Doc("UGcFEEx4kUzS7Y90NogY").Snapshots(ctx)
	for {
		snap, err := it.Next()
		if status.Code(err) == codes.DeadlineExceeded {
			return fmt.Errorf("Snapshots.Next: %w", err)
		}
		if !snap.Exists() {
			fmt.Fprintf(os.Stdout, "Document no longer exists\n")
			return nil
		}

		callback(snap.Data())
	}
}

func writeCode(ctx context.Context, client *firestore.Client, userId string, code string) error {
	_, err := client.Collection("users").Doc(userId).Set(ctx, map[string]interface{}{
		"code": code,
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}
