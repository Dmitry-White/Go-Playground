package main

import (
	"fmt"
	"time"
)

func isWeekend(date time.Time) bool {
	weekday := date.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return true
	}
	return false
}

func nextBusinessDay(date time.Time) time.Time {
	day_duration := 24 * time.Hour

	targetDay := date
	for {
		targetDay = targetDay.Add(day_duration)
		if !isWeekend(targetDay) {
			break
		}
	}

	return targetDay
}

func nextDay(date time.Time) (time.Time, time.Weekday) {
	fmt.Println("Date: ", date, date.Weekday())
	nextDate := nextBusinessDay(date)
	fmt.Println("Next: ", nextDate, nextDate.Weekday())
	return nextDate, nextDate.Weekday()
}
