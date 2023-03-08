package data

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	DB             *sql.DB
	DataSourceName string
	Port           string
	Router         *mux.Router
}
