package services

import (
	"go-applied-concurrency/api/models"
	"sync"
)

type Result struct {
	Latest models.Statistics
	lock   sync.Mutex
}

type IResult interface {
	Get()
	Combine(stats models.Statistics)
}

// Get returns the save statistics result
func (r *Result) Get() models.Statistics {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.Latest
}

// Update safely updates the statistics result
func (r *Result) Combine(stats models.Statistics) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Latest = models.Combine(r.Latest, stats)
}
