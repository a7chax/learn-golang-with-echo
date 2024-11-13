package router

import (
	"database/sql"
	handler "echo-golang/handler/note"
	repository "echo-golang/repository/note"
	service "echo-golang/service/note"
	user_service "echo-golang/service/user"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, db *sql.DB) {
	noteRepo := repository.NoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NoteHandler(noteService)

	r := e.Group("/note")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user_service.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	fmt.Println("Note Router")
	r.GET("", noteHandler.GetNote)
	r.POST("", noteHandler.InsertNote)
	r.DELETE("/:id", noteHandler.DeleteNoteById)
	r.PUT("/:id", noteHandler.UpdateNoteById)
}
