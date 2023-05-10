package main

import (
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

/*
	More Info: https://gobyexample.com/atomic-counters
*/

var (
	total     = expvar.NewInt("total")
	totalSize uint64
)

// ############################## HTTP Server ######################################
func uploadHandler(resw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(resw, err.Error(), http.StatusBadRequest)
	}

	dataSize := len(data)
	atomic.AddUint64(&totalSize, uint64(dataSize))
	total.Add(int64(dataSize))

	fmt.Fprintln(resw, "Uploaded!")
}

func totalSizeHandler(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resw, "size=%d\n", totalSize)
}

func uploadServer() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/totalSize", totalSizeHandler)

	err := http.ListenAndServe(UPLOAD_SERVER_PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// #################################################################################

// ############################## HTTP Client ######################################
func uploadSize() error {
	fmt.Println("Not Implemented")
	return nil
}

// #################################################################################
