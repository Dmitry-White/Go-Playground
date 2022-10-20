package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const REQUEST_TIMEOUT = 30 * time.Millisecond

func getPosts() {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	url := fmt.Sprintf("%s/%s", API["BASE_URL"], API["POSTS"])

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error: can't create a request %s", API["BASE_URL"])
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Error: can't call %s", API["BASE_URL"])
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}

func addPost(post Post) string {
	url := fmt.Sprintf("%s/%s", API["BASE_URL"], API["POSTS"])

	postBytes, err := json.Marshal(post)
	if err != nil {
		log.Fatalf("error: can't marshal - %s", err)
	}

	bytesReader := bytes.NewReader(postBytes)

	response, err := http.Post(url, "application/json", bytesReader)
	if err != nil {
		log.Fatalf("Error: can't call %s, results in %v", API["BASE_URL"], err)
	}

	defer response.Body.Close()

	result, _ := io.ReadAll(response.Body)

	return string(result)
}
