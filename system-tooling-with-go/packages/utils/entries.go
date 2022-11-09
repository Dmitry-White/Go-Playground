package utils

import (
	"bufio"
	"packages/core"
	"strings"
)

func ProcessEntries(r *bufio.Reader) string {
	entry, _ := r.ReadString('\n')
	trimmedEntry := strings.TrimSpace(entry)
	entryName := core.Names[trimmedEntry]
	entryMessage := core.Messages[trimmedEntry]
	entrySays := Greeting(entryName, entryMessage)

	return entrySays
}
