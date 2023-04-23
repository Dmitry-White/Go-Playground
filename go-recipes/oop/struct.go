package main

import (
	"fmt"
	"time"
)

type Event struct {
	id   string
	time time.Time
}

type DoorEvent struct {
	Event
	action string
}

type TermostatEvent struct {
	Event
	temperature float64
}

func getDoorEvent(id string, time time.Time, action string) interface{} {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	doorEvent := DoorEvent{
		Event:  Event{id, time},
		action: action,
	}
	return &doorEvent
}

func getTermostatEvent(id string, time time.Time, temperature float64) interface{} {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	doorEvent := TermostatEvent{
		Event:       Event{id, time},
		temperature: temperature,
	}
	return &doorEvent
}
