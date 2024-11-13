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

//	func restricted(c echo.Context) error {
//		user := c.Get("user").(*jwt.Token)
//		claims := user.Claims.(*user_service.JwtCustomClaims)
//		name := claims.Name
//		return c.String(http.StatusOK, "Welcome "+name+"!")
//	}
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
	// s := middleware.NewStats()
	// e.Use(s.Process)
	router.InitRouter(e, db)
	// echoMiddlare := middleware.NewBasicAuth(userService)
	// e.Use(middleware.BasicAuth(echoMiddlare))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte("your-secret-key"),
	// }))

	e.GET("/user", userHandler.GetAllUser)

	routeLogin := e.Group("/login")
	routeLogin.Use(middleware.BasicAuth())
	routeLogin.POST("", userHandler.LoginUser)

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
