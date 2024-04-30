package storage

type Storage interface {
	Connect() error
	Upload(objectName, filePath string) error
	GetSignedURL(objectName, path string) error
}
