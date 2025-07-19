package chat

import (
	"context"
	"encoding/json"

	"github.com/Adityadangi14/solh_ai/db"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/Adityadangi14/solh_ai/prompt"
	"github.com/Adityadangi14/solh_ai/renderer"
	"google.golang.org/genai"
)

func SendPrompt(ctx context.Context, query string, userId string) (string, error) {
	initializers.AppLogger.Info("Framing prompt", "userId", userId, "query", query)

	response := prompt.Frameprompt(query, userId)

	result, err := initializers.GemClient.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(response),
		nil,
	)

	if err != nil {
		initializers.AppLogger.Error("Gemini content generation failed", "error", err, "userId", userId)
		return "", err
	}

	res := result.Text()
	initializers.AppLogger.Info("Gemini response received", "userId", userId)

	renderedComp, err := renderer.Render(res)
	if err != nil {
		initializers.AppLogger.Error("Rendering failed", "error", err)
		return "", err
	}

	marshaledRes, err := json.Marshal(renderedComp)
	if err != nil {
		initializers.AppLogger.Error("Failed to marshal rendered result", "error", err)
		return "", err
	}

	initializers.AppLogger.Info("Prompt processed successfully", "userId", userId)
	return string(marshaledRes), nil
}

func SaveChatData(prop map[string]any) {
	initializers.AppLogger.Info("Saving chat data", "data", prop)

	_, err := db.SaveData(prop)

	if err != nil {
		initializers.AppLogger.Error("Failed to save chat data", "error", err)
	} else {
		initializers.AppLogger.Info("Chat data saved successfully")
	}
}
