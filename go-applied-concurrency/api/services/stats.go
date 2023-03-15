package services

import (
	"go-applied-concurrency/api/models"
)

type Statistics struct {
	Result    *Result
	Processed <-chan models.Order
	Done      <-chan struct{}
}

type IStatistics interface {
	GetStats() models.Statistics
}

func New(processed <-chan models.Order, done <-chan struct{}) IStatistics {
	s := Statistics{
		Result:    &Result{},
		Processed: processed,
		Done:      done,
	}

	go ProcessStats(&s)

	return &s
}

func (s *Statistics) GetStats() models.Statistics {
	return s.Result.Get()
}

// reconcile is a helper method which saves stats object
// back into the statistics services
func (s *Statistics) reconcile(stats models.Statistics) {
	s.Result.Combine(stats)
}
