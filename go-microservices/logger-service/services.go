package main

import (
	"go-microservices/logger-service/dal"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Services) write(requestPayload *RequestPayload) (*mongo.InsertOneResult, error) {
	logEntry := dal.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	entry, err := s.Models.LogEntry.Insert(logEntry)
	if err != nil {
		log.Println("[write] Model Error: ", err)
		return nil, err
	}

	return entry, nil
}

func (s *Services) all() ([]*dal.LogEntry, error) {
	entries, err := s.Models.LogEntry.GetAll()
	if err != nil {
		log.Println("[all] Model Error: ", err)
		return nil, err
	}

	return entries, nil
}

func (s *Services) writeRPC(requestPayload *RequestPayload) (*mongo.InsertOneResult, error) {
	logEntry := dal.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	entry, err := s.Models.LogEntry.Insert(logEntry)
	if err != nil {
		log.Println("[write] Model Error: ", err)
		return nil, err
	}

	return entry, nil
}
