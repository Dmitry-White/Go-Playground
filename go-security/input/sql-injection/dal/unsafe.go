package dal

import (
	"database/sql"
	_ "embed"
	"fmt"
)

//go:embed dmlHacky.sql
var dmlHackySQL string

func insertHackyLog(db *sql.DB, timestamp, message string) error {
	query := fmt.Sprintf(dmlHackySQL, timestamp, message)

	_, err := db.Exec(query)
	return err
}
