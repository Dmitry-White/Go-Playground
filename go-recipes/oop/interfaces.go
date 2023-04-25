package main

import "fmt"

type Alarm struct {
	id     string
	action string
}

func (a *Alarm) ID() string {
	return a.id
}

func (a *Alarm) Kind() string {
	return "alarm"
}

type Camera struct {
	id  string
	fov int
}

func (c *Camera) ID() string {
	return c.id
}

func (c *Camera) Kind() string {
	return "camera"
}

type Sensor interface {
	ID() string
	Kind() string
}

func printSensors(sensors []Sensor) {
	for _, sensor := range sensors {
		fmt.Printf("%s <%s>\n", sensor.ID(), sensor.Kind())
	}
}

func getSensors() []Sensor {
	entranceCamera := Camera{"Main Entrance", 120}
	smokeAlarm := Alarm{"Office", "smoke"}

	sensors := []Sensor{&entranceCamera, &smokeAlarm}
	printSensors(sensors)

	return sensors
}
