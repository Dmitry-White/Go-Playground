package main

type AppConfig struct {
	PORT int
	ADDR string
}

type Routes struct {
	INDEX  string
	HEALTH string
}
