package dal

import "time"

const RETRY_LIMIT = 10
const RETRY_TIMEOUT = 2 * time.Second
const DB_TIMEOUT = 3 * time.Second

const PASSWORD_COST = 12

var QUERIES = Queries{
	DELETE_BY_ID:    "queries/deleteById.sql",
	GET_ALL:         "queries/getAll.sql",
	GET_BY_EMAIL:    "queries/getByEmail.sql",
	GET_BY_ID:       "queries/getById.sql",
	INSERT:          "queries/insert.sql",
	UPDATE_BY_ID:    "queries/updateById.sql",
	UPDATE_PASSWORD: "queries/updatePassword.sql",
}
