package smtp

import "time"

const RETRY_LIMIT = 10
const RETRY_TIMEOUT = 2 * time.Second

var PATHS = Paths{
	EMAILS:    "./emails",
	TEMPLATES: "./templates",
}
