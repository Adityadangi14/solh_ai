package initializers

import (
	"context"
	"os"

	"google.golang.org/genai"
)

var GemClient *genai.Client

func ConnectToGemini() {
	key := os.Getenv("gemini_api_key")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		AppLogger.Error("Failed to connect to Gemini", "error", err)
		panic("Gemini client initialization failed")
	}

	GemClient = client
	AppLogger.Info("Successfully connected to Gemini client")
}
