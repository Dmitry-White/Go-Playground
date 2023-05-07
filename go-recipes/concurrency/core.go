package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	SEQUENTIAL_STRATEGY = "sequential"
	CONCURRENT_STRATEGY = "concurrent"
)

type OperationsFunc func(time.Time) (float64, error)

func getOperationsFunc(strategy string) (OperationsFunc, error) {
	choice := fmt.Sprintf("Using %s operations", strategy)
	log.Println(choice)

	switch strategy {
	case SEQUENTIAL_STRATEGY:
		return monthDistance, nil
	case CONCURRENT_STRATEGY:
		return monthDistance, nil
	}
	return nil, errors.New("this strategy is not supported")
}
