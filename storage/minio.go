package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio(
	endpoint *string,
	accessKeyID *string,
	secretAccessKey *string,
) (*minio.Client, error) {
	minioClient, err := minio.New(*endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(*accessKeyID, *secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

type MinioStorage struct {
	Client     *minio.Client
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
}

func NewMinioStorage(endpoint *string, accessKey *string, secretKey *string) (*MinioStorage, error) {
	client, err := minio.New(*endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(*accessKey, *secretKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return &MinioStorage{Client: client, Endpoint: *endpoint, AccessKey: *accessKey, SecretKey: *secretKey}, nil
}

func (m *MinioStorage) UploadFile(file *multipart.FileHeader, bucketName, objectName string, contentType string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	_, err = m.Client.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("http://%s/%s/%s", m.Endpoint, bucketName, objectName)
	return fileURL, nil
}

func (m *MinioStorage) DeleteFile(bucketName, objectName string) error {
	err := m.Client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (m *MinioStorage) GetFile(bucketName, objectName string) (string, error) {
	fileURL := fmt.Sprintf("http://%s/%s/%s", m.Endpoint, bucketName, objectName)
	return fileURL, nil
}
