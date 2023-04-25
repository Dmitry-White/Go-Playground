package main

import "fmt"

type Thermostat struct {
	id    string
	value float64
}

func (t *Thermostat) get() float64 {
	return t.value
}

func (t *Thermostat) set(value float64) {
	t.value = value
}

func (*Thermostat) Kind() string {
	return "thermostat"
}

func getThermostat() *Thermostat {
	thermostat := Thermostat{"Living Room", 16.2}

	fmt.Printf("%s before: %.2f\n", thermostat.id, thermostat.get())
	thermostat.set(18)

	return &thermostat
}
