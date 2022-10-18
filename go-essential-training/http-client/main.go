package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getUsers() {
	url := fmt.Sprintf("%s/%s", API["BASE_URL"], API["USERS"])

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: can't call %s", API["BASE_URL"])
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}

func main() {
	getUsers()
}
