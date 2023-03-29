package main

import (
	"fmt"
	"go-security/auth/authentication/middleware"
	"go-security/auth/authentication/services"
	"net/http"
)

func healthHandler(resw http.ResponseWriter, req *http.Request) {
	health, err := services.CheckHealth()
	if err != nil {
		http.Error(resw, "Health check failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(resw, health.Status)
}

func messagesHandler(resw http.ResponseWriter, req *http.Request) {
	userContextKey := middleware.ContextKey("user")

	user, ok := req.Context().Value(userContextKey).(*services.User)
	if !ok {
		http.Error(resw, "No user", http.StatusInternalServerError)
		return
	}

	messages := services.GetMessages(user)
	fmt.Fprintln(resw, messages)
}
