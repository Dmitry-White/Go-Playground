package main

import (
	"encoding/json"
	"go-security/output/sensitive-data/dal"
	"log"
	"net/http"
)

func handleDataSafe(resw http.ResponseWriter, user *dal.User) {
	friends, err := dal.GetUserFriends(user)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(resw, "Can't get friends", http.StatusBadRequest)
		return
	}

	audit := &dal.AuditLog{
		Origin: "handleDataSafe",
		Data:   user,
	}
	audit.Track()

	resw.Header().Set("Content-Type", "application/json")

	jsonWriter := json.NewEncoder(resw)
	jsonWriter.Encode(friends)
}
