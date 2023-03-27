package dal

import "time"

type Message struct {
	Time    time.Time
	From    string
	Content string
}

type Params struct {
	Name     string
	Messages []Message
}
