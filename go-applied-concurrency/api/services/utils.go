package services

import (
	"fmt"
	"go-applied-concurrency/api/models"
	"math/rand"
	"time"
)

func randomSleep() {
	rand.Seed(time.Now().UnixNano())

	duration := time.Duration(rand.Intn(500)) * time.Millisecond

	time.Sleep(duration)
}

// processOrder is a helper function that incorporates
// the current order in the stats service
func processOrder(order *models.Order) models.Statistics {
	// Simulate processing as a costly operation
	randomSleep()

	// Completed orders increment and add to the revenue
	if order.Status == models.OrderStatus_Completed {
		return models.Statistics{
			CompletedOrders: 1,
			Revenue:         *order.Total,
		}
	}

	// Otherwise the order is rejected
	return models.Statistics{
		RejectedOrders: 1,
	}
}

// ProcessStats is the overall processing method that listens to incoming orders
func ProcessStats(s *Statistics) {
	fmt.Println("Stats processing started!")
	for {
		select {
		case order := <-s.Processed:
			processedStats := processOrder(&order)
			s.Stats <- processedStats
			fmt.Printf("Processing Stats %s completed\n", order.ID)
		case <-s.Done:
			fmt.Println("Stats processing stopped!")
			return
		}
	}
}
