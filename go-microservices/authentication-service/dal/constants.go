package dal

import "time"

const RETRY_LIMIT = 10
const RETRY_TIMEOUT = 2 * time.Second
const DB_TIMEOUT = 3 * time.Second

const PASSWORD_COST = 12

var QUERIES = Queries{
	DELETE_BY_ID: "deleteById.sql",
	GET_ALL:      "getAll.sql",
	GET_BY_EMAIL: "getByEmail.sql",
	GET_BY_ID:    "getById.sql",
	INSERT:       "insert.sql",
	UPDATE_BY_ID: "updateById.sql",
}
