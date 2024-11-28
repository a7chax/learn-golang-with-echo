package service

import (
	"echo-golang/storage"
	"mime/multipart"
)

type StorageService struct {
	Storage storage.Storage
}

func NewStorageService(storage storage.Storage) *StorageService {
	return &StorageService{Storage: storage}
}

func (s *StorageService) UploadFile(file *multipart.FileHeader, bucketName, objectName string, contentType string) (string, error) {
	return s.Storage.UploadFile(file, bucketName, objectName, contentType)
}
