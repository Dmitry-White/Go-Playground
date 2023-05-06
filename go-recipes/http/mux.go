package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

const SERVER_MUX_PORT = ":8081"

// Book is information about book
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// isbn -> book
var booksDB = map[string]Book{
	"0062225677": {
		Title:  "The Colour of Magic",
		Author: "Terry Pratchett",
		ISBN:   "0062225677",
	},
	"0765394855": {
		Title:  "Old Man's War",
		Author: "John Scalzi",
		ISBN:   "0765394855",
	},
}

func getBook(isbn string) (Book, error) {
	book, ok := booksDB[isbn]
	if !ok {
		return Book{}, fmt.Errorf("unknown ISBN: %q", isbn)
	}

	return book, nil
}

func handleGetBook(resw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	isbn := vars["isbn"]

	book, err := getBook(isbn)
	if err != nil {
		log.Printf("Error getting from DB: %q", isbn)
		http.Error(resw, err.Error(), http.StatusNotFound)
		return
	}
	log.Printf("Book: %+v", book)

	resw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(resw)

	encoderErr := encoder.Encode(book)
	if encoderErr != nil {
		log.Printf("Error encoding: %s", encoderErr)
		http.Error(resw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServerMux(wg *sync.WaitGroup) {
	router := mux.NewRouter()

	router.HandleFunc("/books/{isbn}", handleGetBook).Methods("GET")

	err := http.ListenAndServe(SERVER_MUX_PORT, router)
	if err != nil {
		wg.Done()
		log.Fatal(err)
	}
}
