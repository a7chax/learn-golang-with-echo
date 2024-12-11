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

	routeUser := e.Group("/user")

	routeUser.GET("", userHandler.GetUser)
	routeUser.GET("/all", userHandler.GetAllUser, middleware.JWT())
	routeUser.POST("/login", userHandler.LoginUser, middleware.BasicAuth())
	routeUser.POST("/refresh", userHandler.RefreshToken, middleware.JWT())
	routeUser.POST("/register", userHandler.RegisterUser, middleware.BasicAuth())
}
