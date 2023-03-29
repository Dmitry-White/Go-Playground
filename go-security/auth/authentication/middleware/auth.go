package middleware

import (
	"context"
	"go-security/auth/authentication/services"
	"net/http"
)

type Auth struct {
	next http.Handler
	user ContextKey
}

func (a *Auth) ServeHTTP(resw http.ResponseWriter, req *http.Request) {
	token := services.GetToken(req)
	user := services.GetUser(token)
	if user == nil {
		http.Error(resw, "Bad Authentication", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(req.Context(), a.user, user)
	req = req.WithContext(ctx)

	a.next.ServeHTTP(resw, req)
}

func WithAuth(handler OriginalHandler) *Auth {
	return &Auth{
		next: http.HandlerFunc(handler),
		user: ContextKey("user"),
	}
}
