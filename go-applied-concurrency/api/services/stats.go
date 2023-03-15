package services

import (
	"fmt"
	"go-applied-concurrency/api/models"
)

type Statistics struct {
	Result    *Result
	Processed <-chan models.Order
	Done      <-chan struct{}
	Stats     chan models.Statistics
}

type IStatistics interface {
	GetStats() models.Statistics
}

const WorkerCount = 3

func New(processed <-chan models.Order, done <-chan struct{}) IStatistics {
	s := Statistics{
		Result:    &Result{},
		Processed: processed,
		Done:      done,
		Stats:     make(chan models.Statistics, WorkerCount),
	}

	for i := 0; i < WorkerCount; i++ {
		go ProcessStats(&s)
	}

	// Reconcile from all goroutines
	go s.reconcile()

	return &s
}

func (s *Statistics) GetStats() models.Statistics {
	return s.Result.Get()
}

// reconcile is a helper method which saves stats object
// back into the statistics services
func (s *Statistics) reconcile() {

	fmt.Println("Reconcile processing started!")
	for {
		select {
		case processedStats := <-s.Stats:
			s.Result.Combine(processedStats)
			fmt.Printf("Processing Reconcile %v completed\n", processedStats)
		case <-s.Done:
			fmt.Println("Reconcile processing stopped!")
			return
		}
	}
}
