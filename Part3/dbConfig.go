package main

import (
	"database/sql"
	"log"
)

var database *sql.DB

func dbConnect() {
	//create connection string
	connStr := "user=postgres password=admin dbname=northwind_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")
}
