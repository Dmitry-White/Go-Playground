package main

type Config struct {
	PORT int
	ADDR string
}

type Routes struct {
	INDEX string
	JS    string
}

type Paths struct {
	JS        string
	PAGES     string
	TEMPLATES string
}
