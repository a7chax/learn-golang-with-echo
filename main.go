package main

import (
	"echo-golang/data/database"
	user_handler "echo-golang/handler/user"
	"echo-golang/middleware"
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
	e.Debug = true

	router.InitRouter(e, db)

	e.GET("/user", userHandler.GetAllUser)

	routeLogin := e.Group("/login")
	routeLogin.POST("", userHandler.LoginUser, middleware.BasicAuth())

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
