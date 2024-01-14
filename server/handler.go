package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"firebase.google.com/go/messaging"
	"log"
	"math/rand"
	"net/http"
)

type AuthPayload struct {
	Code         *string `json:"code"`
	UserId       *string `json:"userid"`
	AndroidToken *string `json:"androidToken"`
}

type CallPayload struct {
	UserId string `json:"userid"`
	Number string `json:"number"`
	Secret string `json:"secret"`
}

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func parseJson(request *http.Request, payload interface{}) error {
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		log.Printf("error decoding json: %s", err)
		return err
	}
	return nil
}

// TODO: refactor as a function that gets: method and http.ResponseWriter
func getHandler(method string, action func(w http.ResponseWriter, request *http.Request)) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "chrome-extension://hnedcaamblfpdpkblfocnjchfbianddf")
		w.Header().Set("Access-Control-Allow-Methods", method)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if request.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if request.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		action(w, request)
	}
}

func getLogoutHandler(ctx context.Context, firestoreClient *firestore.Client, messagingClient *messaging.Client) func(w http.ResponseWriter, request *http.Request) {
	action := func(w http.ResponseWriter, request *http.Request) {
		var callPayload CallPayload
		parseJson(request, &callPayload)
		log.Printf("callPayload: %s", callPayload)
		// TODO: check secret
		user := getUserById(ctx, firestoreClient, callPayload.UserId)
		if user["secret"] == callPayload.Secret {
			deleteUser(ctx, firestoreClient, callPayload.UserId)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
	return getHandler(http.MethodPost, action)
}

func getCodeHandler(ctx context.Context, firestoreClient *firestore.Client, messagingClient *messaging.Client) func(w http.ResponseWriter, request *http.Request) {
	action := func(w http.ResponseWriter, request *http.Request) {
		var codePayload AuthPayload
		parseJson(request, &codePayload)
		log.Printf("codePayload: code: %s", codePayload.Code)
		log.Printf("codePayload: androidToken: %s", codePayload.AndroidToken)
		log.Printf("codePayload: userId: %s", codePayload.UserId)
		if codePayload.AndroidToken != nil {
			// initial request from Android, user sends code and androidToken and user is created
			userId := createUser(ctx, firestoreClient, *codePayload.Code, *codePayload.AndroidToken)
			log.Printf("userCreaed id: %s", userId)
			w.WriteHeader(http.StatusCreated)
		} else {
			// matching request from android, user sends code and chrome is sent secret and userid
			userId := getUserIdByCode(ctx, firestoreClient, *codePayload.Code)
			androidToken := getAndroidToken(ctx, firestoreClient, userId)
			secret := RandString()
			saveUserSecret(ctx, firestoreClient, userId, secret)
			message := &messaging.Message{
				Data: map[string]string{
					"ok":     "true",
					"secret": secret,
				},
				Token: androidToken,
			}
			log.Printf("androidToken: %s", androidToken)
			err := sendMessage(ctx, messagingClient, message)

			if err {
				payloady := map[string]string{
					"secret": secret,
					"userid": userId,
				}
				json.NewEncoder(w).Encode(payloady)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
	return getHandler(http.MethodPost, action)
}

func getCallHandler(ctx context.Context, firestoreClient *firestore.Client, messagingClient *messaging.Client) func(w http.ResponseWriter, request *http.Request) {
	action := func(w http.ResponseWriter, request *http.Request) {
		var callPayload CallPayload
		parseJson(request, &callPayload)
		log.Printf("callPayload: %s", callPayload)
		// TODO: check secret
		user := getUserById(ctx, firestoreClient, callPayload.UserId)
		if user["secret"] == callPayload.Secret {
			log.Printf("user authenticated: %s", user)
			androidToken := user["androidToken"].(string)
			log.Printf("androidToken: %s", androidToken)
			message := &messaging.Message{
				Data: map[string]string{
					"number": callPayload.Number,
				},
				Token: androidToken,
				Notification: &messaging.Notification{
					Title: "click to call",
					Body:  callPayload.Number,
				},
			}
			sendMessage(ctx, messagingClient, message)
		}
	}
	return getHandler(http.MethodPost, action)
}

func HealthCheckHandler(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
}
