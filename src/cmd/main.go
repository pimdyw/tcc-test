package main

import (
	"database/sql"
	"net/http"
	"tcc-test/config"
	"tcc-test/internal/repository"
	"tcc-test/internal/transport"
	"tcc-test/internal/usecase"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	userTable := config.GetPostgresDB()
	initEcho(userTable)
	initEcho(userTable).Logger.Fatal(initEcho(userTable).Start(":8080"))
}

func initEcho(userTable *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	userRepository := repository.NewUserRepository(userTable)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userTransport := transport.NewUserTransport(userUsecase)

	e.GET("/health-check", healthCheck)
	e.POST("/user", userTransport.CreateUser)
	e.GET("/users", userTransport.GetUsers)
	e.GET("/user", userTransport.GetUser)

	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(200, "success")
}
