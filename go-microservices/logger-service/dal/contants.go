package dal

import "time"

const RETRY_LIMIT = 10
const RETRY_TIMEOUT = 2 * time.Second
const DB_TIMEOUT = 3 * time.Second
