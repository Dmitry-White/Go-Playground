package main

import (
	"database/sql"
	"go-security/input/sql-injection/dal"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)

	connStr := "user=postgres password=s3cr3t sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Content: ", content)
	log.Println("DB: ", db)

	dal.CreateTable(db)

	dal.InsertLog("hacky")(db, time.Now(), content)
}
