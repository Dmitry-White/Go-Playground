package dal

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

//go:embed queries/*
var queriesFolder embed.FS

var retryCounts int64

// New is the function used to create an instance of the dal package. It returns the type
// Models, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

func Connect() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("DB is not yet ready...")
			retryCounts++
		} else {
			log.Println("Connected to DB!")
			return connection
		}

		if retryCounts > RETRY_LIMIT {
			log.Println(err)
			return nil
		}

		log.Printf("Backing off for %d seconds...\n", RETRY_TIMEOUT/time.Second)
		time.Sleep(RETRY_TIMEOUT)
		continue
	}
}
