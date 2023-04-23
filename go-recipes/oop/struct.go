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
