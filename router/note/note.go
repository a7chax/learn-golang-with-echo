package router

import (
	"database/sql"
	handler "echo-golang/handler/note"
	user_handler "echo-golang/handler/user"
	"echo-golang/middleware"
	repository "echo-golang/repository/note"
	user_repository "echo-golang/repository/user"
	service "echo-golang/service/note"
	user_service "echo-golang/service/user"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err := cv.validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		errors := make(map[string]string)
// 		for _, err := range err.(validator.ValidationErrors) {
// 			// Custom error messages for each field and validation rule
// 			switch err.Field() {
// 			case "Name":
// 				if err.Tag() == "required" {
// 					errors["name"] = "Name is required"
// 				} else if err.Tag() == "min" {
// 					errors["name"] = "Name must be at least 9 characters"
// 				} else if err.Tag() == "max" {
// 					errors["name"] = "Name must be at most 10 characters"
// 				}
// 			case "Email":
// 				if err.Tag() == "required" {
// 					errors["email"] = "Email is required"
// 				} else if err.Tag() == "email" {
// 					errors["email"] = "Invalid email format"
// 				}
// 			}
// 		}
// 		return echo.NewHTTPError(http.StatusBadRequest, errors)
// 	}
// 	return nil
// }

func InitRouter(e *echo.Echo, db *sql.DB) {
	noteRepo := repository.NoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NoteHandler(noteService)

	userRepo := user_repository.UserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := user_handler.UserHandler(userService)

	routeLogin := e.Group("/login")
	routeLogin.POST("", userHandler.LoginUser, middleware.BasicAuth())

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
