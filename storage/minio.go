package storage

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/barelyhuman/go/env"
	"github.com/minio/minio-go"
)

type S3Storage struct {
	client *minio.Client
	bucket string
}

func NewAWSStorage(bucket string) *S3Storage {
	clientId := env.Get("STORAGE_CLIENT_ID", "")
	clientSecret := env.Get("STORAGE_CLIENT_SECRET", "")
	endpoint := env.Get("STORAGE_ENDPOINT", "")
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, clientId, clientSecret, ssl)
	if err != nil {
		log.Fatal(err)
	}

	bucketExists, _ := client.BucketExists(bucket)

	if !bucketExists {
		err := client.MakeBucket(
			bucket,
			"blr1",
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// List all Spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces {
		fmt.Println(space.Name)
	}

	return &S3Storage{
		client: client,
		bucket: bucket,
	}
}

func (a *S3Storage) Connect() error {
	return nil
}

func (a *S3Storage) Upload(objectName string, data bytes.Buffer) error {
	dataBytes := bytes.NewReader(data.Bytes())
	_, err := a.client.PutObject(
		a.bucket,
		objectName,
		dataBytes,
		dataBytes.Size(),
		minio.PutObjectOptions{},
	)
	return err
}

func (a *S3Storage) HasObject(objectName string) bool {
	obj, err := a.client.StatObject(a.bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return len(obj.Key) > 0
}

func (a *S3Storage) GetSignedURL(objectName string) (string, error) {
	url, err := a.client.PresignedGetObject(
		a.bucket, objectName, time.Minute*15, url.Values{},
	)
	return url.String(), err
}

func (a *S3Storage) ListObjects() []MicroObject {
	doneCh := make(chan struct{})
	defer close(doneCh)
	recursive := true
	collection := []MicroObject{}
	for obj := range a.client.ListObjectsV2(a.bucket, "", recursive, doneCh) {
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
	err := a.client.RemoveObject(a.bucket, objectName)
	if err != nil {
		return false, err
	}
	return true, nil
}
