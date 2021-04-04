package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // only get driver sql
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "devrian:@tcp(localhost:3306)/library?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
