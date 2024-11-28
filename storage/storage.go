package storage

import "mime/multipart"

type Storage interface {
	UploadFile(file *multipart.FileHeader, bucketName, objectName string, contentType string) (string, error)
	DeleteFile(bucketName, objectName string) error
	GetFile(bucketName, objectName string) (string, error)
}
