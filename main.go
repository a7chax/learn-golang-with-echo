package main

import (
	"echo-golang/data/database"
	router_note "echo-golang/router/note"
	router_storage "echo-golang/router/storage"
	router_user "echo-golang/router/user"
	"echo-golang/storage"
	"fmt"
	"io"

	"os"

	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ptr(s string) *string {
	return &s
}

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := database.ConnectDatabaseNote(
		ptr(os.Getenv("DB_HOST")),
		ptr(os.Getenv("DB_USER")),
		ptr(os.Getenv("DB_PASSWORD")),
		ptr(os.Getenv("DB_NAME")),
		ptr(os.Getenv("DB_SSL_MODE")),
	)

	minioStorage, _ := storage.NewMinioStorage(
		ptr(os.Getenv("MINIO_ENDPOINT")),
		ptr(os.Getenv("MINIO_ACCESS_KEY")),
		ptr(os.Getenv("MINIO_SECRET_KEY")),
	)

	// minioClient, _ := storage.ConnectMinio(
	// 	ptr(os.Getenv("MINIO_ENDPOINT")),
	// 	ptr(os.Getenv("MINIO_ACCESS_KEY")),
	// 	ptr(os.Getenv("MINIO_SECRET_KEY")),
	// )

	e := echo.New()
	e.Debug = true

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Static("/", "public")
	// e.POST("/upload", upload)

	router_note.InitNoteRouter(e, db)
	router_user.InitUserRouter(e, db)
	router_storage.InitStorageRouter(e, minioStorage)

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
