package initializers

import (
	"context"
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to gemini client")
	GemClient = client
}
