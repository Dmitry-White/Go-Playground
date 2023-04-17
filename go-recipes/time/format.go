package main

import (
	"fmt"
	"time"
)

// Formatting converts time object ot string
func getBirthday(date time.Time) time.Duration {
	fmt.Println("Default: ", date) // 1940-10-09 18:30:00 +0000 UTC

	fmt.Println("Custom 1: ", date.Format("2006-01-02"))    // 1940-10-09
	fmt.Println("Custom 2: ", date.Format("Mon, Jan 02"))   // Wed, Oct 09
	fmt.Println("RFC3339: ", date.Format(time.RFC3339Nano)) // 1940-10-09T18:30:00Z

	return 3500 * time.Millisecond
}
