package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	router := NewRouter()
	dbConnect()
	log.Fatal(http.ListenAndServe(":8888", router))
}
