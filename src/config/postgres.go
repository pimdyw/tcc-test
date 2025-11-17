package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))

	fmt.Println(err)
	if err != nil {
		fmt.Println("sql error", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return err
	}

	fmt.Println("Successfully connected to database")

	// Ensure the table exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        address VARCHAR(255) NOT NULL,
        birth_date TIMESTAMP NULL DEFAULT NULL,
        age BIGINT NOT NULL
    )`)
	if err != nil {
		return err
	}
	return err
}

func GetPostgresDB() *sql.DB {
	return db
}
