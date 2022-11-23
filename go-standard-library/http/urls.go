package main

import (
	"fmt"
	"net/url"
)

func urls() {
	s := "https://www.example.com:8000/user?username=dmitrywhite"

	// TODO: Use the Parse function to parse the URL content
	parsedUrl, _ := url.Parse(s)
	fmt.Println("URL:", parsedUrl)
	fmt.Printf("URL struct: %+v\n", *parsedUrl)
	fmt.Println("Port: ", parsedUrl.Port())

	// TODO: Extract the query components into a Values struct
	values := parsedUrl.Query()
	fmt.Printf("Query map: %+v\n", values)
	fmt.Println(values["username"])

	// TODO: Create a URL from components
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	fmt.Println("Original:", newUrl.String())

	newUrl.Host = "temp.com"

	fmt.Println("Host Updated:", newUrl.String())

	// TODO: Create a new Values struct and encode it as a query string
	newValues := &url.Values{
		"x": []string{"100"},
		"y": []string{"200", "300"},
	}

	newValues.Add("z", "somestr")

	newUrl.RawQuery = newValues.Encode()

	fmt.Println("Query Updated:", newUrl.String())
}
