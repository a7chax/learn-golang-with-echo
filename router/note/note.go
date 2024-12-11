package router_note

import (
	"database/sql"
	handler "echo-golang/handler/note"
	"echo-golang/middleware"
	repository "echo-golang/repository/note"
	user_repository "echo-golang/repository/user"

	service "echo-golang/service/note"

	"github.com/labstack/echo/v4"
)

func InitNoteRouter(e *echo.Echo, db *sql.DB) {
	noteRepo := repository.NoteRepository(db)
	userRepo := user_repository.UserRepository(db)
	noteService := service.NewNoteService(noteRepo, userRepo)
	noteHandler := handler.NoteHandler(noteService)

	routeNote := e.Group("/note")

	routeNote.Use(middleware.JWT())
	routeNote.GET("", noteHandler.GetNote)
	routeNote.POST("", noteHandler.InsertNote)
	routeNote.DELETE("/:id", noteHandler.DeleteNoteById)
	routeNote.PUT("/:id", noteHandler.UpdateNoteById)
}
