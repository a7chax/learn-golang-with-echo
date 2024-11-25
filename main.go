package main

import (
	"echo-golang/data/database"
	router_note "echo-golang/router/note"
	router_user "echo-golang/router/user"

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

	e := echo.New()
	e.Debug = true

	router_note.InitNoteRouter(e, db)
	router_user.InitUserRouter(e, db)

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
