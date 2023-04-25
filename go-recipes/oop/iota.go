package main

import "fmt"

type LogLevel uint8

/*
In some programming languages (C++, Go, etc.),
iota is used to represent and generate
an array of consecutive integers.

https://en.wikipedia.org/wiki/Iota

https://github.com/golang/go/wiki/Iota
*/
const (
	Debug LogLevel = iota + 1
	Warning
	Info
	Error
)

func initLogLevel() error {
	fmt.Println("Not Implemented")
	return nil
}
