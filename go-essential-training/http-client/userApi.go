package main

import (
	"bytes"
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

func addUser(user User) map[string]interface{} {
	url := fmt.Sprintf("%s/%s", API["BASE_URL"], API["USERS"])

	// Unlike regular variable declarations,
	// a short variable declaration may redeclare variables
	// provided they were originally declared earlier in the same block with the same type,
	// and at least one of the non-blank variables is new.
	// As a consequence, redeclaration can only appear in a multi-variable short declaration.
	// Redeclaration does not introduce a new variable; it just assigns a new value to the original.
	//
	// Therefore, there are multiple "err" variables seemingly "redeclared"
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("error: can't marshal - %s", err)
	}

	bytesReader := bytes.NewReader(userBytes)

	response, err := http.Post(url, "application/json", bytesReader)
	if err != nil {
		log.Fatalf("Error: can't call %s, results in %v", API["BASE_URL"], err)
	}

	defer response.Body.Close()

	buffer := make([]byte, 9999)
	length, _ := response.Body.Read(buffer)

	var data map[string]interface{}
	jsonUnmarshalErr := json.Unmarshal(buffer[:length], &data)
	if jsonUnmarshalErr != nil {
		log.Fatalf("Error: can't unmarshal %v", jsonUnmarshalErr)
	}

	return data
}
