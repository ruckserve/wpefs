package gcloud_api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

type GC_client struct {
	ctx         context.Context
	client      storage.Client
	bktHandle   storage.BucketHandle
}

func (gc *GC_client) PutObj(key string, bytes []byte) (err error) {
	var wc storage.Writer
	defer wc.Close()
	wc = gc.bktHandle.Object(key).NewWriter(gc.ctx)

	if _, err := wc.Write(bytes); err != nil {
		log.Fatalf("Failed to write object %s to cloud storage", key)
		// TODO: retry probably
		return nil, err
	}
	return len(bytes), nil
}

func (gc *GC_client) GetObj(key string) (bytes []byte, err error) {
	var rc storage.Reader
	rc, err = gc.bktHandle.Object(key).NewReader(gc.ctx)
	if (err != nil) {
		log.Fatalf("Failed to get object for %s", key)
	}
	defer rc.Close()

	for {
		size := rc.Size()
		bytes = make([]byte, size)
		_, err := rc.Read(bytes)
		if n := rc.Remain(); n < 0 || err != nil {
			continue
		}
		return
	}
}
//
//func main() {
//	// Creates a client.
//	client, err := storage.NewClient(ctx)
//	if err != nil {
//		log.Fatalf("Failed to create client: %v", err)
//	}
//
//	// Sets the name for the new bucket.
//	bucketName := "hackathon-docker-storage-driver"
//
//	// Creates a Bucket instance.
//	bucket := client.Bucket(bucketName)
//	filename := "filename1"
//
//	writer := bucket.Object(filename).NewWriter(ctx)
//	if _, err := writer.Write([]byte("hello world")); err != nil {
//		log.Fatalf("Failed to write to object %v", filename)
//	}
//	if err := writer.Close(); err != nil {
//		log.Fatalf("Failed to close object %v", filename)
//	}
//
//	fmt.Printf("Wrote hello world to %v in %v.\n", filename, bucketName)
//}
