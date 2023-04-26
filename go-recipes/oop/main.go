package main

import (
	"fmt"
	"go-recipes/oop/misc"
	"time"
)

func main() {
	fmt.Println("------------- Struct -------------")
	fmt.Printf("%+v\n", getDoorEvent("front door", time.Now(), "open"))
	fmt.Printf("%+v\n", getTermostatEvent("front door", time.Now(), 25))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Methods ------------")
	fmt.Printf("%+v\n", getThermostat())
	fmt.Println("----------------------------------")

	fmt.Println("----------- Interfaces -----------")
	fmt.Printf("%#v\n", getSensors())
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Empty -------------")
	fmt.Println(getCounts())
	fmt.Println("----------------------------------")

	fmt.Println("-------------- IOTA --------------")
	fmt.Println(initLogLevel(Debug))
	fmt.Println(initLogLevel(Warning))
	fmt.Println(initLogLevel(Info))
	fmt.Println(initLogLevel(Error))
	fmt.Println(initLogLevel(19))
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Image -------------")
	fmt.Println(misc.GetFace())
	fmt.Println("----------------------------------")
}
