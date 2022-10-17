package main

// Generally, encoding/decoding JSON refers to the process
// of actually reading/writing the character data to a string or binary form.
// Marshaling/Unmarshaling refers to the process
// of mapping JSON types from and to Go data types and primitives.

// Encode => Stream
// Marshal => String
func main() {
	decodeRequest()
	encodeResponse(1000)
	unmarshalRequest()
	marshalResponse(1000)
}
