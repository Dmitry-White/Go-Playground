package main

import (
	"fmt"
	"time"
)

func nextBusinessDay(time.Time) time.Time {
	fmt.Println("Not Implemented")
	return time.Now()
}

func nextDay(date time.Time) (time.Time, time.Weekday) {
	fmt.Println("Date: ", date, date.Weekday())
	nextDate := nextBusinessDay(date)
	fmt.Println("Next: ", nextDate, nextDate.Weekday())
	return nextDate, nextDate.Weekday()
}
