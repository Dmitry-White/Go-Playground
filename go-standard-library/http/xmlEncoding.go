package main

import (
	"encoding/xml"
	"fmt"
)

func xmlEncoding() {
	// TODO: Declare some sample data
	people := []Person{
		{"John Doe", "123 Anywhere street", 36, nil},
		{"Jane Doe", "456 Elsewhere Blvd", 31, []string{"Purple", "Yellow"}},
	}

	// TODO: The MarshalIndent function indents the XML text
	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result:\n%s\n%s\n", xml.Header, string(result))
}
