package main

import (
	"fmt"
	"net/http"
)

func checkLink(channel chan bool, link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- false
		return false
	}
	fmt.Println(link, "is up!")
	channel <- true
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

	channel := make(chan bool)

	for _, link := range links {
		go checkLink(channel, link)
	}

	for i := 0; i < len(links); i++ {
		fmt.Printf("Status is %v\n", <-channel)
	}
}
