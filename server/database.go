package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
)

func sendMessage(ctx context.Context, client *messaging.Client, message *messaging.Message) bool {
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println("error sending message:", err)
		return false
	} else {
		fmt.Println("Successfully sent message:", response)
		return true
	}
}

func createUser(ctx context.Context, client *firestore.Client, code string, androidToken string) *string {
	log.Printf("code: %s, androidToken: %s", code, androidToken)
	docRef, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"code":         code,
		"androidToken": androidToken,
	})

	if err == nil {
		return &docRef.ID
	}
	return nil
}

func saveUserSecret(ctx context.Context, firestoreClient *firestore.Client, userId string, secret string) {
	firestoreClient.Collection("users").Doc(userId).Set(ctx, map[string]interface{}{"secret": secret}, firestore.MergeAll)
}

func deleteUser(ctx context.Context, firestoreClient *firestore.Client, userId string) {
	firestoreClient.Collection("users").Doc(userId).Delete(ctx)
}

func getUserById(ctx context.Context, firestoreClient *firestore.Client, userId string) map[string]interface{} {
	doc, err := firestoreClient.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		log.Printf("error getting user: %s", err)
		return nil
	}
	var data map[string]interface{}
	doc.DataTo(&data)
	return data
}

func getAndroidToken(ctx context.Context, firestoreClient *firestore.Client, userId string) string {
	user := getUserById(ctx, firestoreClient, userId)
	log.Printf("user: %s", user)
	return user["androidToken"].(string)
}

func getUserIdByCode(ctx context.Context, firestoreClient *firestore.Client, code string) string {
	doc, err := firestoreClient.Collection("users").Where("code", "==", code).Documents(ctx).Next()
	if err != nil {
		log.Printf("error getting user by code: %s", err)
		return ""
	}
	return doc.Ref.ID
}
