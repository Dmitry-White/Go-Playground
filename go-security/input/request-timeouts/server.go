package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(resw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	start := time.Now()
	n, err := io.Copy(io.Discard, req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d bytes in %v", n, time.Since(start))

	payload := fmt.Sprintf("%d bytes digested", n)
	fmt.Fprintln(resw, payload)
}

func getSafeServer() (*http.Server, error) {
	return &http.Server{
		Addr:              ":8080",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}, nil
}

func getHackyServer() (*http.Server, error) {
	return &http.Server{
		Addr: ":8080",
	}, nil
}

func getClient(strategy string) func() (*http.Server, error) {
	choice := fmt.Sprintf("Using %s server", strategy)
	log.Println(choice)

	return func() (*http.Server, error) {
		switch strategy {
		case "safe":
			return getSafeServer()
		case "hacky":
			return getHackyServer()
		}
		return nil, errors.New("this strategy is not supported")
	}
}
