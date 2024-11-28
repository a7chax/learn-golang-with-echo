package router_storage

import (
	handler "echo-golang/handler/storage"
	service "echo-golang/service/storage"
	"echo-golang/storage"

	"github.com/labstack/echo/v4"
)

func InitStorageRouter(e *echo.Echo, storage storage.Storage) {
	storageService := service.NewStorageService(storage)
	storageHandler := handler.NewStorageHandler(storageService)
	routeStorage := e.Group("/storage")
	routeStorage.POST("", storageHandler.UploadFile)
}
