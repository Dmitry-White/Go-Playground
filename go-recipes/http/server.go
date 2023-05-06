package main

import (
	"encoding/json"
	"hash/fnv"
	"io"
	"log"
	"net/http"
)

// DB is a database
type DB struct{}

// Add adds a metric to the database
func (db *DB) Add(m interface{}) (uint32, error) {
	strMetric, err := json.Marshal(m)
	if err != nil {
		return 0, err
	}

	h := fnv.New32()
	h.Write(strMetric)

	return h.Sum32(), nil
}

var db *DB

const PORT = ":8080"

func handleMetric(resw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(resw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer req.Body.Close()

	log.Println("Request:", req)

	m := Metric{}
	reader := io.LimitReader(req.Body, SIZE)
	decoder := json.NewDecoder(reader)

	decoderErr := decoder.Decode(&m)
	if decoderErr != nil {
		log.Printf("Error decoding: %s", decoderErr)
		http.Error(resw, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.Add(m)
	if err != nil {
		log.Printf("Error adding to DB: %s", err)
		http.Error(resw, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Metric(id=%d): %+v", id, m)

	resw.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"id": id,
	}
	encoder := json.NewEncoder(resw)

	encoderErr := encoder.Encode(data)
	if encoderErr != nil {
		log.Printf("Error encoding: %s", encoderErr)
		http.Error(resw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServer() interface{} {
	http.HandleFunc("/metric", handleMetric)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	return nil
}
