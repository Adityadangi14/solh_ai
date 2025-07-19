package initializers

import (
	"fmt"
	"log/slog"
	"os"
)

var AppLogger *slog.Logger

func InitAppLogs() {
	logDir := "logs"
	logFilePath := logDir + "/app.log"

	// Ensure logs/ directory exists
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		panic(fmt.Sprintf("error creating log directory: %v", err))
	}

	// Open the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file: %v", err))
	}

	// Create JSON handler for structured logging
	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	// Set global logger
	logger := slog.New(handler)
	slog.SetDefault(logger)
	AppLogger = logger
}
