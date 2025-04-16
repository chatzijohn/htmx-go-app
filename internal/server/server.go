package server

import (
	"my-app/internal/config"
	"my-app/internal/server/routes"
	"my-app/internal/service"
	"net/http"
)

func New(serverConfig *config.ServerConfig, services *service.ServiceContainer) http.Handler {
	mux := http.NewServeMux()

	// Serve static files (CSS, JS, etc.)
	staticFs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", staticFs))

	// Register application routes from central router
	appRouter := routes.NewRouter(services)
	appRouter.RegisterRoutes(mux)

	return mux
}
