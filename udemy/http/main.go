package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

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
	// data := make([]byte, 99999)
	// resp.Body.Read(data)

	// log.Println(string(data))

	io.Copy(os.Stdout, resp.Body)
}
