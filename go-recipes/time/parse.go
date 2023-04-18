package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	TIME_STRATEGY     = "time"
	DURATION_STRATEGY = "duration"
)

type ParserFunc = func(string) (interface{}, error)

func parseTime(string) (interface{}, error) {
	return time.Now(), nil
}

func parseDuration(string) (interface{}, error) {
	return time.Minute, nil
}

func getParser(strategy string) (ParserFunc, error) {
	choice := fmt.Sprintf("Using %s parser", strategy)
	log.Println(choice)

	switch strategy {
	case TIME_STRATEGY:
		return parseTime, nil
	case DURATION_STRATEGY:
		return parseDuration, nil
	}
	return nil, errors.New("this strategy is not supported")
}

// Parsing converts string to time object
func fromString(strategy string, datetime string) interface{} {
	parser, err := getParser(strategy)
	if err != nil {
		log.Fatalln(err)
	}

	date, err := parser(datetime)
	if err != nil {
		return err
	}

	return date
}
