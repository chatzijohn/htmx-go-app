package server

import (
	"my-app/internal/config"
	"net/http"
)

func New(serverConfig *config.ServerConfig) http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	return handler
}
