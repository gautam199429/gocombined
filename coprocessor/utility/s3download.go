package utilities

import (
	"fmt"
	"time"

	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	schema        string
	schemaChanged bool
)

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
				fmt.Println("Error updating schema:", err)
				continue
			}
			oldHash := HashSHA256(schema)
			newHash := HashSHA256(newSchema)
			if oldHash == newHash {
				fmt.Println("Schema has not changed.")
				schemaChanged = false
				continue
			} else {
				fmt.Println("Schema has changed.")
				schema = newSchema
				schemaChanged = true
				fmt.Println("Schema updated at", time.Now())
			}
		}
	}()
}

func GetSchema() (string, bool, error) {
	if schema == "" {
		newSchema, err := DownloadSchemaAsString()
		if err != nil {
			return "", false, fmt.Errorf("failed to read object content: %w", err)
		}
		schema = newSchema
		return schema, true, nil
	} else {
		return schema, schemaChanged, nil
	}
}
