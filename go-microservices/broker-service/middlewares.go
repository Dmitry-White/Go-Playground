package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var loggerMiddleware = middleware.Logger

var healthMiddleware = middleware.Heartbeat(ROUTES.HEALTH)

var corsOptions = cors.Options{
	AllowedOrigins:   []string{},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300,
}
var corsMiddleware = cors.Handler(corsOptions)
