package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

func UploadFile(w io.Writer, bucket, object string, file multipart.File) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	f := file

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucket).Object(object)

	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", object)
	return nil
}

func GetSignedUrl(bucket, object string) string {
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)
	url, err := client.Bucket(bucket).SignedURL(object, &storage.SignedURLOptions{})
	if err != nil {
		return fmt.Errorf("storage.SignedURL: %w", err).Error()
	}
	return url
}
