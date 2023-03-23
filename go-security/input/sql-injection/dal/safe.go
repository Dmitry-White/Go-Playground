package dal

import (
	"database/sql"
	_ "embed"
)

//go:embed dmlSafe.sql
var dmlSafeSQL string

func insertSafeLog(db *sql.DB, timestamp string, message string) error {
	_, err := db.Exec(dmlSafeSQL, timestamp, message)
	return err
}
