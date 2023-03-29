package middleware

import "net/http"

type OriginalHandler = func(resw http.ResponseWriter, req *http.Request)

type ContextKey string
