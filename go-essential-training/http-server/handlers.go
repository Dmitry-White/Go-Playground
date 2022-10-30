package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
	defer r.Body.Close()

	requestDecoder := json.NewDecoder(r.Body)
	requestContent := &MathRequest{}

	err := requestDecoder.Decode(requestContent)
	if err != nil {
		log.Printf("Error: bad JSON %s", err)
		http.Error(w, "Bad JSON", http.StatusBadRequest)
	}

	isMathOperation := strings.Contains("+-*/", requestContent.Op)
	if !isMathOperation {
		log.Printf("Error: bad operator %q", requestContent.Op)
		http.Error(w, "Unknown Operator", http.StatusBadRequest)
	}

	responseContent := &MathResponse{}
	switch requestContent.Op {
	case "+":
		responseContent.Result = requestContent.Left + requestContent.Right
	case "-":
		responseContent.Result = requestContent.Left - requestContent.Right
	case "*":
		responseContent.Result = requestContent.Left * requestContent.Right
	case "/":
		if requestContent.Right == 0.0 {
			responseContent.Error = "Division by 0"
		} else {
			responseContent.Result = requestContent.Left / requestContent.Right
		}
	default:
		responseContent.Error = fmt.Sprintf("Unknown Operation: %s", requestContent.Op)
	}

	w.Header().Set("Content-Type", "application/json")
	if responseContent.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	responseEncoder := json.NewEncoder(w)
	encoderError := responseEncoder.Encode(responseContent)
	if encoderError != nil {
		log.Printf("Error: can't encode %v - %s", responseContent, err)
	}
}
