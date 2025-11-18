package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"tcc-test/internal/repository"
	"tcc-test/internal/transport"
	"tcc-test/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
	if err != nil {
		log.Println("sql error", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("ping error", err.Error())
		return
	}

	log.Println("Successfully connected to database")

	// Ensure the table exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		surname VARCHAR(255) NOT NULL,
        address TEXT NOT NULL,
        birth_date TIMESTAMP NULL DEFAULT NULL,
        age BIGINT NOT NULL
    )`)
	if err != nil {
		return
	}

	initEcho(db)
	initEcho(db).Logger.Fatal(initEcho(db).Start(":8080"))
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
	e.GET("/user/:id", userTransport.GetUser)

	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(200, "success")
}
