package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"name"`
	Address    string   `json:"addr"`
	Age        int      `json:"-"`
	FaveColors []string `json:"favs,omitempty"`
}

func jsonEncoding() {
	// TODO: Create some people data
	people := []person{
		{"John Doe", "123 Anywhere street", 36, nil},
		{"Jane Doe", "456 Elsewhere Blvd", 31, []string{"Purple", "Yellow"}},
	}

	// TODO: Marshal is used to convert a data structure to JSON format
	result, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %s\n", result)
	fmt.Println("Result: ", string(result))

	// TODO: MarshalIndent is used to format the JSON string with indentation
	data, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data: %s\n", data)
	fmt.Println("Data: ", string(data))
}
