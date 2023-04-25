package main

import (
	"log"
)

type ClickEvent struct{}
type HoverEvent struct{}

var eventsCounter = make(map[string]int)

func record(event interface{}) {
	switch event.(type) {
	case *ClickEvent:
		eventsCounter["click"]++
	case *HoverEvent:
		eventsCounter["hover"]++
	default:
		log.Printf("warning: unknown event: %#v of type %T\n", event, event)
	}
}

func getCounts() map[string]int {
	record(&ClickEvent{})
	record(&HoverEvent{})
	record(&ClickEvent{})
	record(3)
	return eventsCounter
}
