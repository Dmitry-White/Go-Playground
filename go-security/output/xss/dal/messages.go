package dal

import (
	"bytes"
	"fmt"
	"time"
)

func LoadMessages(user string) ([]Message, error) {
	ms := []Message{
		{
			time.Date(2021, time.September, 23, 10, 33, 17, 0, time.UTC),
			"Pippin",
			"What about second breakfast?",
		},
		{
			time.Date(2021, time.September, 23, 14, 15, 32, 0, time.UTC),
			"Samy",
			"Where's the ring? <script>alert('Pwned!')</script>",
		},
	}

	return ms, nil
}

func FormateMessages(ms []Message) string {
	var buf bytes.Buffer
	for _, m := range ms {
		ts := m.Time.Format("2006-01-02T15:04")
		fmt.Fprintf(&buf, "<p>[%s %s] %s</p><hr />\n", ts, m.From, m.Content)
	}
	return buf.String()
}
