package main

import (
	"fmt"
	"time"
)

func parseDate(datetime string) (time.Time, error) {
	duration, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return time.Time{}, err
	}

	return duration, nil
}

func convertFromTo(datetime time.Time, fromTZ, toTZ string) (time.Time, time.Time, error) {
	from, err := time.LoadLocation(fromTZ)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	fromTime := datetime.In(from)

	to, err := time.LoadLocation(toTZ)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	toTime := fromTime.In(to)

	return fromTime, toTime, nil
}

func tzConvert(datetime string) interface{} {
	date, err := parseDate(datetime)
	if err != nil {
		return err
	}

	fromTime, toTime, err := convertFromTo(date, "America/Chicago", "America/New_York")
	if err != nil {
		return err
	}

	fmt.Printf("Begin: %v - %v\n", date.Location(), date.String())
	fmt.Printf("Between: %v - %v\n", fromTime.Location(), fromTime.String())
	fmt.Printf("End: %v - %v\n", toTime.Location(), toTime.String())

	return toTime
}
