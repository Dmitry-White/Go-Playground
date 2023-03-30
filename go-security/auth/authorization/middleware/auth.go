package middleware

import (
	"net/http"
)

type Access struct {
	next http.Handler
	acl  ACL
}

func (a *Access) checkAccess(req *http.Request, whitelist *Whitelist) bool {
	if contains(whitelist.Methods, req.Method) &&
		contains(whitelist.Resources, req.URL.Path) {
		return true
	}

	return false
}

func (a *Access) ServeHTTP(resw http.ResponseWriter, req *http.Request) {
	user := req.URL.Query().Get("user")
	userWhitelist := a.acl[user]

	if !a.checkAccess(req, &userWhitelist) {
		http.Error(resw, "You don't have access", http.StatusUnauthorized)
		return
	}

	a.next.ServeHTTP(resw, req)
}

func WithAccess(handler OriginalHandler) *Access {
	return &Access{
		next: http.HandlerFunc(handler),
		acl: ACL{
			"joe": {
				Methods:   []string{"GET", "POST"},
				Resources: []string{"/profile"},
			},
			"jane": {
				Methods:   []string{"GET"},
				Resources: []string{"/profile"},
			},
		},
	}
}
