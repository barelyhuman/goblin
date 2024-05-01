package storage

import "bytes"

type Storage interface {
	Connect() error
	HasObject(objectName string) bool
	Upload(objectName string, data bytes.Buffer) error
	GetSignedURL(objectName string) (string, error)
}
