package main

import (
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", *resp)

	// Allocating memory beforehand
	// since Reader.Read function can't automatically
	// resize the slice if it's already full
	// therefore if passed an empty slice
	// it won't be able to fill it with data
	//
	// data := make([]byte, 99999)
	// resp.Body.Read(data)
	// log.Println(string(data))

	// Built-in way of outputing Response
	// Works sort of like piping ReadStream to WriteStream in Node.js
	// io.Copy(os.Stdout, resp.Body)

	logger := logWriter{}
	io.Copy(logger, resp.Body)
}

func (logWriter) Write(data []byte) (int, error) {
	log.Println(string(data))
	log.Println("Just wrote this many bytes: ", len(data))

	return len(data), nil
}
