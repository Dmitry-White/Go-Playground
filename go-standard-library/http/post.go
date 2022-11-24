package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func post() {
	const httpbin = "https://httpbin.org/post"

	// TODO: POST operation using Post
	reqBody := strings.NewReader(`
	{
		"field1": "This is field 1",
		"field2": 250
	}
	`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Result: ", string(content))

	// TODO: POST operation using PostForm
	data := url.Values{
		"x": []string{"100"},
		"y": []string{"200", "300"},
	}
	respForm, err := http.PostForm(httpbin, data)
	if err != nil {
		panic(err)
	}

	defer respForm.Body.Close()

	contentForm, _ := ioutil.ReadAll(respForm.Body)
	fmt.Println("Result: ", string(contentForm))
}
