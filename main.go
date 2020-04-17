package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/dlsniper/dockerdemo/handlers"
)

func main() {

	dataSource := "postgres://goland:goland@%s:5432/goland?sslmode=disable"
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	db, err := sqlx.Connect("postgres", fmt.Sprintf(dataSource, dbHost))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.Home(db))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

