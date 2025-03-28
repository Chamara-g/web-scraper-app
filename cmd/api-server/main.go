package main

import (
	"net/http"

	"github.com/gihanc.dev/web-scraper-app/internal/logger"
	"github.com/gihanc.dev/web-scraper-app/internal/router"
)

// CORS Middleware function to add CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's an OPTIONS request, just return 200 OK (preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {

	router := router.InitializeRoutes()
	router.Use(corsMiddleware)

	logger.Logger.Info("Server starting on port 8080", "version", "1.0")

	if err := http.ListenAndServe(":8080", router); err != nil {
		
		logger.Logger.Error(err.Error(), "error", "response")
	}
}