package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "hackathon-docker-storage-driver"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)
	filename := "filename1"

	writer := bucket.Object(filename).NewWriter(ctx)
	writer.ContentType = "text/plain"
	if _, err := writer.Write([]byte("hello world")); err != nil {
		log.Fatalf("Failed to write to object %v", filename)
	}
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to close object %v", filename)
	}

	fmt.Printf("Wrote hello world to %v in %v.\n", filename, bucketName)
}
