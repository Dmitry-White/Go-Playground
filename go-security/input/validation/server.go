package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	handle, err := getHandleFunc(SAFE_STRATEGY)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := handle(decoder)
	if err != nil {
		errMessage := fmt.Sprintf("Validation errors:\n%v", err)
		log.Println(errMessage)
		http.Error(res, errMessage, http.StatusBadRequest)
		return
	}

	log.Println("Payment:", data)

	fmt.Fprintln(res, data)
}
