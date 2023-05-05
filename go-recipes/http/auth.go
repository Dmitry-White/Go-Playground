package main

import "fmt"

type Credentials struct {
	username string
	password string
}

func basicAuth(url string, credentials Credentials) (interface{}, error) {
	fmt.Println("Not Implemented")
	return nil, nil
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
