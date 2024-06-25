package storage

import (
	"bytes"
	"time"
)

type MicroObject struct {
	LastModified time.Time
	Key          string
}

type Storage interface {
	Connect() error
	HasObject(objectName string) bool
	Upload(objectName string, data bytes.Buffer) error
	GetSignedURL(objectName string) (string, error)
	ListObjects() []MicroObject
	RemoveObject(objectName string) (bool, error)
}
