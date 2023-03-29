package services

import (
	"net/http"
	"strings"
)

type User struct {
	Login string
}

func GetToken(req *http.Request) string {
	// Dummy Implementation: generate JWT, Oauth2 ...
	header := req.Header.Get("Authorization")
	token := strings.TrimPrefix(header, "Bearer ")

	return token
}

func GetUser(token string) *User {
	// Dummy Implementation: check JWT, Oauth2 ...
	if token == "crypto" {
		return &User{"Bob"}
	}
	return nil
}
