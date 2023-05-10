package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	OPERATIONS_SERVER_PORT = ":8080"
	UPLOAD_SERVER_PORT     = ":8081"
	SEQUENTIAL_STRATEGY    = "sequential"
	CONCURRENT_STRATEGY    = "concurrent"
)

type OperationsFunc func(time.Time) (float64, error)

func getOperationsFunc(strategy string) (OperationsFunc, error) {
	choice := fmt.Sprintf("Using %s operations", strategy)
	log.Println(choice)

	switch strategy {
	case SEQUENTIAL_STRATEGY:
		return monthDistanceSequential, nil
	case CONCURRENT_STRATEGY:
		return monthDistanceConcurrent, nil
	}
	return nil, errors.New("this strategy is not supported")
}
