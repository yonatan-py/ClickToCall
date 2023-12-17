package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"time"
)

func sendMessage(ctx context.Context, client messaging.Client, registrationToken string) {
	message := &messaging.Message{
		Data: map[string]string{
			"score": "foooooo",
			"time":  "bar",
		},
		Token: registrationToken,
		Notification: &messaging.Notification{
			Title: "$GOOG up 1.43% on the day",
			Body:  "$GOOG gained 11.80 points to close at 835.67, up 1.43% on the day.",
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

	for true {
		println("sending message")
		sendMessage(ctx, *client, registrationToken)
		time.Sleep(5 * time.Second)
	}
}
