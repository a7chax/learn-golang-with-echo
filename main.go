package main

import (
	"echo-golang/data/database"
	user_handler "echo-golang/handler/user"
	router "echo-golang/router/note"

	user_repository "echo-golang/repository/user"
	user_service "echo-golang/service/user"
	"os"

	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ptr(s string) *string {
	return &s
}

// type (
// 	CustomValidator struct {
// 		validator *validator.Validate
// 	}
// )

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err := cv.validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		errors := make(map[string]string)
// 		for _, err := range err.(validator.ValidationErrors) {
// 			switch err.Field() {
// 			case "Username":
// 				if err.Tag() == "required" {
// 					errors["username"] = "Username is required"
// 				} else if err.Tag() == "min" {
// 					errors["username"] = "Username must be at least 1 character"
// 				} else if err.Tag() == "max" {
// 					errors["username"] = "Username must be at most 100 characters"
// 				}
// 			case "Password":
// 				if err.Tag() == "required" {
// 					errors["password"] = "Password is required"
// 				} else if err.Tag() == "min" {
// 					errors["password"] = "Password must be at least 4 characters"
// 				} else if err.Tag() == "max" {
// 					errors["password"] = "Password must be at most 100 characters"
// 				}
// 			}
// 		}
// 		return echo.NewHTTPError(http.StatusBadRequest, errors)
// 	}
// 	return nil
// }

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

	userRepo := user_repository.UserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := user_handler.UserHandler(userService)

	e := echo.New()

	// e.Validator = &CustomValidator{validator: validator.New()}
	e.Debug = true

	router.InitRouter(e, db)

	e.GET("/user", userHandler.GetAllUser)

	// routeLogin := e.Group("/login")
	// routeLogin.POST("", userHandler.LoginUser, middleware.BasicAuth())

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
