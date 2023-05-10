package main

import (
	"bytes"
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

/*
	More Info: https://gobyexample.com/atomic-counters
*/

var (
	total     = expvar.NewInt("total")
	totalSize uint64
)

const (
	UPLOAD_BASE_URL = "http://localhost:8081"
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
func request(verb, endpoint string) interface{} {
	url := fmt.Sprintf("%s/%s", UPLOAD_BASE_URL, endpoint)
	var resp *http.Response
	var err error

	switch verb {
	case "GET":
		resp, err = http.Get(url)
	case "POST":
		data := []byte(time.Now().String())
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return string(result)
}

func uploadSize() []interface{} {
	uploadResult := request("POST", "upload")
	totalResult := request("GET", "totalSize")
	debugResult := request("GET", "debug/vars")

	return []interface{}{uploadResult, totalResult, debugResult}
}

// #################################################################################
