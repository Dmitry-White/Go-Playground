package data

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}
