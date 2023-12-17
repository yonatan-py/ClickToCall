package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

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

	message := &messaging.Message{
		Data: map[string]string{
			"score": "6dsdsa",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", response)
}
