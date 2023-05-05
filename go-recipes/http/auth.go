package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Credentials struct {
	username string
	password string
}

func basicAuth(url string, credentials Credentials) (interface{}, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(credentials.username, credentials.password)
	fmt.Println("Request:", request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d %s", response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response)

	result := map[string]interface{}{}
	reader, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(reader, &result)
	if err != nil {
		return nil, jsonErr
	}
	fmt.Printf("Result: %#v\n", result)

	return &result, nil
}

func authenticate() interface{} {
	credentials := Credentials{"joe", "baz00ka"}
	url := fmt.Sprintf("%s/%s/%s/%s",
		BASE_URL,
		AUTH_ENDPOINT,
		credentials.username,
		credentials.password,
	)

	resp, err := basicAuth(url, credentials)
	if err != nil {
		return err
	}
	return resp
}
