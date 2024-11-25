package router_note

import (
	"database/sql"
	handler "echo-golang/handler/note"
	repository "echo-golang/repository/note"
	service "echo-golang/service/note"
	user_service "echo-golang/service/user"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitNoteRouter(e *echo.Echo, db *sql.DB) {
	noteRepo := repository.NoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NoteHandler(noteService)

	routeNote := e.Group("/note")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user_service.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	routeNote.Use(echojwt.WithConfig(config))
	routeNote.GET("", noteHandler.GetNote)
	routeNote.POST("", noteHandler.InsertNote)
	routeNote.DELETE("/:id", noteHandler.DeleteNoteById)
	routeNote.PUT("/:id", noteHandler.UpdateNoteById)
}
