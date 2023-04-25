package main

import "fmt"

type Alarm struct {
	id     string
	action string
}

type Camera struct {
	id  string
	fov int
}

type Sensor interface {
	id() string
	kind() string
}

func printSensors(sensors []Sensor) {
	for _, sensor := range sensors {
		fmt.Printf("%s <%s>\n", sensor.id(), sensor.kind())
	}
}

func getSensors() []Sensor {
	entranceCamera := Camera{"Main Entrance", 120}
	smokeAlarm := Alarm{"Office", "smoke"}

	sensors := []Sensor{&entranceCamera, &smokeAlarm}
	printSensors(sensors)

	return sensors
}
