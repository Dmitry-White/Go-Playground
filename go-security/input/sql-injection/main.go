package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	filename := string(data)

	connStr := "user=postgres password=s3cr3t sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Filename: ", filename)
	log.Println("DB: ", db)
}
