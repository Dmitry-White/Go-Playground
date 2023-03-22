package dal

import (
	"database/sql"
	_ "embed"
	"log"
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
