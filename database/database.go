package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Aich@2302@tcp(127.0.0.1:3306)/BetDB")
	if err != nil {
		panic(err.Error())
		return nil
	} else {
		return db
	}
}
