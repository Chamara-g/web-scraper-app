package main

import (
	"log/slog"
	"net/http"
	"os"

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
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	router := router.InitializeRoutes()
	router.Use(corsMiddleware)

	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(router))

	log.Info("server starting on port 8080")

	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("Failed to start server: ", err)
	}
}