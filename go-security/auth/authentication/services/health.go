package services

import (
	"errors"
	"math/rand"
	"time"
)

type Health struct {
	Status string
}

func CheckHealth() (*Health, error) {
	// Dummy Implementation
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	status := rand.Intn(2)

	if status == 1 {
		return nil, errors.New("Unhealthy")
	}

	return &Health{"Healthy"}, nil
}
