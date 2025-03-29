package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	// Create a log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// Create a multi-handler (logs to console and file)
	fileHandler := slog.New(slog.NewJSONHandler(file, nil))
	consoleHandler := slog.New(slog.NewTextHandler(os.Stdout, nil))

	Logger = consoleHandler
	
	// Combine both handlers
	Logger.Info("Logger initialized", "output", "console & file")
	fileHandler.Info("Logger initialized", "output", "file only")
	
}