package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get() {
	const httpbin = "https://httpbin.org/get"

	// TODO: Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		panic(err)
	}

	// TODO: The caller is responsible for closing the response
	defer resp.Body.Close()

	// TODO: We can access parts of the response to get information
	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Protocol:", resp.Proto)
	fmt.Println("Content Length:", resp.ContentLength)

	// TODO: Use a string Builder to build the content from bytes
	var sb strings.Builder

	// TODO: Read the content and write it to the Builder
	content, _ := ioutil.ReadAll(resp.Body)

	bytecount, _ := sb.Write(content)

	// TODO: Format the output
	fmt.Printf("Result (%d): %s\n", bytecount, sb.String())
}
