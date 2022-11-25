package main

import (
	"encoding/json"
	"fmt"
)

func jsonDecoding() {
	// TODO: Declare some sample JSON data to decode
	data := []byte(`
		{
			"name": "John Q Public",
			"addr": "987 Main St",
			"age": 45,
			"favs": ["White", "Gold"]
		}
	`)

	// TODO: JSON will be decoded into a person struct
	var p person

	// TODO: Test to see if the JSON is valid and the decode
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Println("Result:", p)
	}

	// TODO: Data can also be decoded into a map structure
	var unknownFormat map[string]interface{}

	json.Unmarshal(data, &unknownFormat)
	fmt.Println("Result:", unknownFormat)

	for k, v := range unknownFormat {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
