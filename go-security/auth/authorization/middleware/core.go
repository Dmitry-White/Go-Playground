package middleware

import "net/http"

type OriginalHandler func(resw http.ResponseWriter, req *http.Request)

type Whitelist struct {
	Methods   []string
	Resources []string
}

type ACL map[string]Whitelist
