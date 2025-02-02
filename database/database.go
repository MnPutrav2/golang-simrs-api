package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "_:_@tcp(_)/_")
	if err != nil {
		panic(err.Error())
	}

	return db
}
