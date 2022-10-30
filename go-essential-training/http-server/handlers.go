package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything's shiny, Captain. Not to fret.")
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Everything's shiny, Captain. Not to fret.\n"))
}

func handleMath(w http.ResponseWriter, r *http.Request) {
	log.Println("Not implemented")
}
