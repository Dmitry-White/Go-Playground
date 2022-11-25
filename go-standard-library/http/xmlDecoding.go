package main

import (
	"encoding/xml"
	"fmt"

	"github.com/clbanning/mxj/v2"
)

func xmlDecoding() {
	// TODO: Create some sample XML data
	data := []byte(`
		<Person age="31">
			<firstname>Jane Doe</firstname>
			<address>456 Elsewhere Blvd</address>       
			<favecolors>Purple</favecolors>
			<favecolors>Yellow</favecolors>
		</Person>
	`)

	// TODO: Data will be decoded into a person struct
	var p Person

	// TODO: Use the Unmarshal function to decode raw XML data
	xml.Unmarshal(data, &p)
	fmt.Println("Result:", p)

	// TODO: Data can also be decoded into a map structure
	unknownFormat, err := mxj.NewMapXml(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", unknownFormat)
	fmt.Println(
		unknownFormat["Person"].(map[string]interface{})["firstname"],
	)
}
