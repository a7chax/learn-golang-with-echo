package router

import (
	"database/sql"
	user_handler "echo-golang/handler/user"
	repository "echo-golang/repository/user"
	service "echo-golang/service/user"

	"github.com/labstack/echo/v4"
)

func InitRouteUser(e *echo.Echo, db *sql.DB) {
	userRepo := repository.UserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user_handler.UserHandler(userService)

	r := e.Group("/user")
}
