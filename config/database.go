package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connection() {
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3307)/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Database connected ....")
	DB = db
}
