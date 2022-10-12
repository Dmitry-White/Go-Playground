package main

import "fmt"

func makeRequest(link string) bool {
	fmt.Println(link)
	return true
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		makeRequest(link)
	}
}
