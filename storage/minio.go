package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/barelyhuman/go/env"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client       *minio.Client
	bucket       string
	bucketPrefix string
	ctx          context.Context
}

func NewAWSStorage(bucket string) *S3Storage {
	clientId := env.Get("STORAGE_CLIENT_ID", "")
	clientSecret := env.Get("STORAGE_CLIENT_SECRET", "")
	endpoint := env.Get("STORAGE_ENDPOINT", "")
	bucketPrefix := env.Get("STORAGE_BUCKET_PREFIX", "")
	ssl := true
	ctx := context.Background()

	// Initiate a client using DigitalOcean Spaces.
	creds := credentials.NewStaticV4(clientId, clientSecret, "")
	opts := minio.Options{
		Secure: ssl,
		Creds:  creds,
		Region: "blr1",
	}
	client, err := minio.New(endpoint, &opts)
	if err != nil {
		log.Fatal(err)
	}

	bucketExists, _ := client.BucketExists(ctx, bucket)

	fmt.Printf("bucketExists: %v\n", bucketExists)

	if !bucketExists {
		err := client.MakeBucket(
			ctx,
			bucket,
			minio.MakeBucketOptions{
				Region: "blr1",
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// List all Spaces.
	spaces, err := client.ListBuckets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces {
		fmt.Println(space.Name)
	}

	return &S3Storage{
		client:       client,
		bucketPrefix: bucketPrefix,
		bucket:       bucket,
		ctx:          ctx,
	}
}

func (a *S3Storage) Connect() error {
	return nil
}

func (a *S3Storage) Upload(objectName string, data bytes.Buffer) error {
	dataBytes := bytes.NewReader(data.Bytes())
	_, err := a.client.PutObject(
		a.ctx,
		a.bucket,
		filepath.Join(a.bucketPrefix, objectName),
		dataBytes,
		dataBytes.Size(),
		minio.PutObjectOptions{},
	)
	return err
}

func (a *S3Storage) HasObject(objectName string) bool {
	objectKey := objectName

	if !strings.HasPrefix(objectName, a.bucketPrefix) {
		objectKey = filepath.Join(a.bucketPrefix, objectName)
	}

	obj, err := a.client.StatObject(a.ctx, a.bucket, objectKey, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return len(obj.Key) > 0
}

func (a *S3Storage) GetSignedURL(objectName string) (string, error) {
	objectKey := objectName

	if !strings.HasPrefix(objectName, a.bucketPrefix) {
		objectKey = filepath.Join(a.bucketPrefix, objectName)
	}

	url, err := a.client.PresignedGetObject(
		a.ctx,
		a.bucket, objectKey, time.Minute*15, url.Values{},
	)
	return url.String(), err
}

func (a *S3Storage) ListObjects() []MicroObject {
	doneCh := make(chan struct{})
	defer close(doneCh)
	recursive := true
	collection := []MicroObject{}

	for obj := range a.client.ListObjects(a.ctx, a.bucket, minio.ListObjectsOptions{
		Recursive:    recursive,
		WithMetadata: true,
		UseV1:        true,
		Prefix:       filepath.Join(a.bucketPrefix),
	}) {
		if obj.Err != nil {
			fmt.Printf("obj.Err: %v\n", obj.Err)
			continue
		}
		if filepath.Clean(obj.Key) == a.bucketPrefix {
			continue
		}

		collection = append(collection,
			MicroObject{
				LastModified: obj.LastModified,
				Key:          obj.Key,
			},
		)
	}
	return collection
}

func (a *S3Storage) RemoveObject(objectName string) (bool, error) {
	objectKey := objectName

	if !strings.HasPrefix(objectName, a.bucketPrefix) {
		objectKey = filepath.Join(a.bucketPrefix, objectName)
	}

	err := a.client.RemoveObject(a.ctx, a.bucket, objectKey, minio.RemoveObjectOptions{})
	if err != nil {
		return false, err
	}
	return true, nil
}
