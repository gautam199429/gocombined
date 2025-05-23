package utilities

import (
	"fmt"
	"time"

	"bytes"
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var schema string

func DownloadSchemaAsString() (string, error) {
	minioClient, err := minio.New("https://objectstore.e2enetworks.net", &minio.Options{
		Creds:  credentials.NewStaticV4("RQ8AU8QHQRYJXJVVXXXX", "W89811ZSJ86B3J2QTHS031PZECIIJGKWAIQAXXXX", ""),
		Secure: false,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Get the object
	object, err := minioClient.GetObject(context.Background(), "carsauda", "schema.graphql", minio.GetObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get object: %w", err)
	}
	defer object.Close()

	// Read object content into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, object)
	if err != nil {
		return "", fmt.Errorf("failed to read object content: %w", err)
	}

	return buf.String(), nil
}

func StartSchemaUpdater() {
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			newSchema, err := DownloadSchemaAsString()
			if err != nil {
				log.Println("Error updating schema:", err)
				continue
			}
			schema = newSchema // <-- fix: assign to the global variable
			log.Println("Schema updated at", time.Now())
		}
	}()
}
