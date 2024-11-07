package main

import (
	"echo-golang/data/database"
	handler "echo-golang/handler/note"
	repository "echo-golang/repository/note"
	"os"

	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ptr(s string) *string {
	return &s
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

	noteRepo := repository.NoteRepository(db)
	noteHandler := handler.NoteHandler(noteRepo)

	e := echo.New()
	e.GET("/", noteHandler.GetNote)
	e.POST("/insert", noteHandler.InsertNote)
	e.DELETE("/delete/:id", noteHandler.DeleteNoteById)
	e.PUT("/update/:id", noteHandler.UpdateNoteById)
	e.PUT("/update/:id", noteHandler.UpdateNoteById)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
