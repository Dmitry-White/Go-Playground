package main

import "fmt"

type Sensor interface {
	id() string
	kind() string
}

func printSensors(sensors []Sensor) {
	for _, sensor := range sensors {
		fmt.Printf("%s <%s>\n", sensor.id(), sensor.kind())
	}
}

func getSensors() Sensor {
	fmt.Println("Not Implemented")
	return nil
}
