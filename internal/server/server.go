package server

import (
	"my-app/internal/config"
	"my-app/internal/server/routes"
	"my-app/internal/service"
	"net/http"
)

func New(config *config.AppConfig, services *service.ServiceContainer) http.Handler {
	mux := http.NewServeMux()

	// Register application routes from central router
	appRouter := routes.NewRouter(services)
	appRouter.RegisterRoutes(mux)

	if config.ENVIRONMENT == "development" {
		// Serve static files (CSS, JS, etc.)
		staticFs := http.FileServer(http.Dir("./web/static"))
		mux.Handle("/static/", http.StripPrefix("/static", staticFs))
	}

	return mux
}
