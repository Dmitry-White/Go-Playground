package main

import (
	"encoding/json"
	"go-security/output/sensitive-data/dal"
	"net/http"
)

func handleDataHacky(resw http.ResponseWriter, user *dal.User) {
	friends, err := dal.GetUserFriends(user)
	if err != nil {
		http.Error(resw, err.Error(), http.StatusNotFound)
	}

	audit := &dal.AuditLog{
		Origin: "handleDataHacky",
		Data:   user,
	}
	audit.Track()

	resw.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(resw)
	jsonWriter.Encode(friends)
}
