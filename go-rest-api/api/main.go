package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DB: ", db)
}
