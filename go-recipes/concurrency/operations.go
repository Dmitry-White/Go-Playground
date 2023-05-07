package main

import (
	"fmt"
	"log"
	"time"
)

func monthDistance(time.Time) (float64, error) {
	return 0, nil
}

func operations(strategy string) func(month time.Time) interface{} {
	operationsFunc, err := getOperationsFunc(strategy)
	if err != nil {
		log.Fatalln(err)
	}

	return func(month time.Time) interface{} {
		start := time.Now()

		distance, err := operationsFunc(month)
		if err != nil {
			log.Fatal(err)
		}

		duration := time.Since(start)

		result := fmt.Sprintf("distance=%.2f, duration=%v\n", distance, duration)

		return result
	}
}
