package main

import (
	"echo-golang/data/database"
	handler "echo-golang/handler/note"
	repository "echo-golang/repository/note"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	db, _ := database.ConnectDatabaseNote()
	noteRepo := repository.NoteRepository(db)
	noteHandler := handler.NoteHandler(noteRepo)

	e := echo.New()
	e.GET("/", noteHandler.GetNote)
	e.POST("/insert", noteHandler.InsertNote)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
