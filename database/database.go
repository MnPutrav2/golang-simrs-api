package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	godotenv.Load()

	dbQuery := os.Getenv("DB_CONN")

	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		panic(err.Error())
	}

	return db
}
