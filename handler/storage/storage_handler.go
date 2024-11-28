package handler

import (
	service "echo-golang/service/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StorageHandler struct {
	StorageService *service.StorageService
}

func NewStorageHandler(storageService *service.StorageService) *StorageHandler {
	return &StorageHandler{StorageService: storageService}
}

func (h *StorageHandler) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file"})
	}

	bucketName := "testbucket"
	objectName := file.Filename
	contentType := file.Header.Get("Content-Type")

	url, err := h.StorageService.UploadFile(file, bucketName, objectName, contentType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"url": url})
}
