package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8000"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", helloWorld)

	fmt.Printf("Server started and listening on localhost%s", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
