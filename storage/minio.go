package storage

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	Client     *minio.Client
	BucketName string
}

func (s *MinioStorage) Connect() error {
	endpoint := os.Getenv("MINIO_URL")
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			os.Getenv("MINIO_ROOT_USER"),
			os.Getenv("MINIO_ROOT_PASSWORD"),
			""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}

	s.Client = client
	return nil
}

func (s *MinioStorage) Upload(objectName string, filePath string) error {

	ctx := context.Background()
	bucketName := s.BucketName

	info, err := s.Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	fmt.Println(info)
	return nil
}

func (s *MinioStorage) GetSignedURL(objectName string, alias string) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+alias+"\"")
	presignedURL, err := s.Client.PresignedGetObject(context.Background(), s.BucketName, objectName, time.Minute*5, reqParams)
	if err != nil {
		return "", err
	}

	url := os.Getenv("MINIO_URL_PREFIX") + presignedURL.RequestURI()

	return url, nil
}
