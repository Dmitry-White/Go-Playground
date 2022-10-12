package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return false
	}
	fmt.Println(link, "is up!")
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
		checkLink(link)
	}
}
