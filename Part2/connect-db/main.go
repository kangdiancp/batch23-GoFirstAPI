package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func main() {
	//create connection string
	connStr := "user=postgres dbname=northwind_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")
}
