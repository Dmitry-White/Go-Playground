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

func (ll LogLevel) String() string {
	switch ll {
	case Debug:
		return "Debug"
	case Warning:
		return "Warning"
	case Info:
		return "Info"
	case Error:
		return "Error"
	}

	return fmt.Sprintf("Unknown log level: %d", ll)
}

func initLogLevel(level LogLevel) LogLevel {
	logLevel := LogLevel(level)
	return logLevel
}
