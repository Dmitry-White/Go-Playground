package main

import (
	"fmt"
	"go-security/auth/authorization/services"
	"net/http"
)

func handler(resw http.ResponseWriter, req *http.Request) {
	user := req.URL.Query().Get("user")

	switch req.Method {
	case http.MethodGet:
		profile := services.GetProfile(user)
		fmt.Fprintln(resw, profile)

	case http.MethodPost:
		profile := services.UpdateProfile(user)
		fmt.Fprintln(resw, profile)

	case http.MethodOptions:
		resw.Header().Set("Allow", "GET, POST, OPTIONS")
		resw.WriteHeader(http.StatusNoContent)

	default:
		resw.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(resw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
