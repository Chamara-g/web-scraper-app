package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gihanc.dev/web-scraper-app/internal/logger"
	"github.com/gihanc.dev/web-scraper-app/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	router := router.InitializeRoutes()
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(router))

	log.Info("server starting on port 8080")

	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("Failed to start server: ", err)
	}
}