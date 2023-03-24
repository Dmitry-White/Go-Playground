package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getSafeReader(req *http.Request) (io.Reader, error) {
	return io.LimitReader(req.Body, 1), nil
}

func getHackyReader(req *http.Request) (io.ReadCloser, error) {
	return req.Body, nil
}

func getReader(strategy string) func(req *http.Request) (io.Reader, error) {
	choice := fmt.Sprintf("Using %s reader", strategy)
	log.Println(choice)

	return func(req *http.Request) (io.Reader, error) {
		switch strategy {
		case "safe":
			return getSafeReader(req)
		case "hacky":
			return getHackyReader(req)
		}
		return nil, errors.New("this strategy is not supported")
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	reader, err := getReader("hacky")(req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(reader)

	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalln(err)
	}

	n := len(data)

	payload := fmt.Sprintf("%d bytes stored", n)
	fmt.Fprintln(res, payload)
}
