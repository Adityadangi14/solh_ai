package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		AppLogger.Error("Failed to load .env file", "error", err)
		panic("Error loading .env file")
	}

	AppLogger.Info(".env file loaded successfully")
}
