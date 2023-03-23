package dal

import (
	"database/sql"
	_ "embed"
	"errors"
	"log"
	"time"
)

/*
	Go source files that import "embed" can use the //go:embed directive
	to initialize a variable of type string, []byte, or FS
	with the contents of files read from the package directory or subdirectories
	at compile time.
*/

//go:embed ddl.sql
var ddlSQL string

func CreateTable(db *sql.DB) error {
	log.Println("ddlSQL: ", ddlSQL)
	_, err := db.Exec(ddlSQL)
	return err
}

func InsertLog(strategy string) func(db *sql.DB, time time.Time, message string) error {
	return func(db *sql.DB, time time.Time, message string) error {
		timestamp := time.Format("2006-01-02 15:04:05") // Format time for SQL
		log.Println("Timestamp:", timestamp)

		switch strategy {
		case "safe":
			return insertSafeLog(db, timestamp, message)
		case "hacky":
			return insertHackyLog(db, timestamp, message)
		}
		return errors.New("this strategy is not supported")
	}
}
