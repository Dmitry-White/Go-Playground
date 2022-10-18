package main

import (
	"encoding/json"
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

func getUser(id string) User {
	url := fmt.Sprintf("%s/%s/%s", API["BASE_URL"], API["USERS"], id)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: can't call %s, results in %v", API["BASE_URL"], err)
	}

	defer response.Body.Close()

	buffer := make([]byte, 9999)
	length, _ := response.Body.Read(buffer)

	var user User
	jsonErr := json.Unmarshal(buffer[:length], &user)
	if jsonErr != nil {
		log.Fatalf("Error: can't unmarshal %v", jsonErr)
	}

	return user
}

func main() {
	getUsers()

	user := getUser("1")
	log.Printf("User: %+v", user)
}
