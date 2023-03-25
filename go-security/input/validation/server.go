package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Payment struct {
	Time   string
	User   string
	To     string
	Amount int
}

func handler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	var p Payment
	err := decoder.Decode(&p)
	if err != nil {
		log.Print(err)
		http.Error(res, "Bad JSON", http.StatusBadRequest)
		return
	}

	log.Println("Payment:", p)

	fmt.Fprintln(res, p)
}
