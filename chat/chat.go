package chat

import (
	"context"
	"log"

	"github.com/Adityadangi14/solh_ai/db"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/Adityadangi14/solh_ai/prompt"
	"google.golang.org/genai"
)

func SendPrompt(ctx context.Context, query string, userId string) (string, error) {

	response := prompt.Frameprompt(query, userId)

	result, err := initializers.GemClient.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(response),
		nil,
	)

	if err != nil {
		return "", err
	}
	res := result.Text()

	return res, nil
}

func SaveChatData(prop map[string]any) {

	_, err := db.SaveData(prop)

	if err != nil {
		log.Fatalln("Failed to save chat data")
	}
}
