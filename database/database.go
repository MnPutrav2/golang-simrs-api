package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", "%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
