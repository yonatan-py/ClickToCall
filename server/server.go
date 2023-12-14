package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// listenDocument listens to a single document.
func listenDocument(ctx context.Context, w io.Writer, projectID, collection string) error {
	// projectID := "project-id"
	// Сontext with timeout stops listening to changes.
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("firestore.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Collection(collection).Doc("UGcFEEx4kUzS7Y90NogY").Snapshots(ctx)
	for {
		snap, err := it.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Snapshots.Next: %w", err)
		}
		if !snap.Exists() {
			fmt.Fprintf(w, "Document no longer exists\n")
			return nil
		}
		fmt.Fprintf(w, "Received document snapshot: %v\n", snap.Data())
	}
}

func main() {
	ctx := context.Background()
	projectID := "click-to-call-d2769"
	collection := "users"

	err := listenDocument(ctx, os.Stdout, projectID, collection)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
