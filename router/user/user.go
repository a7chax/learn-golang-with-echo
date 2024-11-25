package router_user

import (
	"database/sql"
	user_handler "echo-golang/handler/user"
	"echo-golang/middleware"
	repository "echo-golang/repository/user"
	service "echo-golang/service/user"

	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, db *sql.DB) {
	userRepo := repository.UserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user_handler.UserHandler(userService)

	routeLogin := e.Group("/user")
	routeLogin.GET("", userHandler.GetAllUser)
	routeLogin.POST("/login", userHandler.LoginUser, middleware.BasicAuth())
}
