package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Category struct {
	CategoryId   int
	CategoryName string
	Description  string
}

var database *sql.DB

func main() {
	//create connection string
	connStr := "user=postgres password=admin dbname=northwind_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")

	http.HandleFunc("/", showCategory)
	http.HandleFunc("/category", displayCategory)

	log.Println("Starting server o 8888")

	errHttp := http.ListenAndServe(":8888", nil)
	if errHttp != nil {
		log.Println(errHttp)
	}
}

func displayCategory(w http.ResponseWriter, r *http.Request) {
	categories := []Category{}

	rows, err := database.Query(`
		select category_id, category_name,
		description from categories
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//tampilkan output row category
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.Description)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
		//load category.html for manipulate with data from slice
	}

	t, _ := template.ParseFiles("category.html")

	t.Execute(w, categories)
}

func showCategory(w http.ResponseWriter, r *http.Request) {
	category := Category{}

	rows, err := database.Query(`
		select category_id, category_name,
		description from categories
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//tampilkan output row category
	for rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.Description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d %s %s\n", category.CategoryId, category.CategoryName, category.Description)
	}
}
