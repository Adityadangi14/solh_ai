package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"github.com/Adityadangi14/solh_ai/db"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/Adityadangi14/solh_ai/prompt"
	"github.com/Adityadangi14/solh_ai/renderer"
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

	fmt.Println(res)

	renderedComp, err := renderer.Render(res)

	if err != nil {
		return "", err
	}

	marshaledRes, err := json.Marshal(renderedComp)

	if err != nil {
		return "", err
	}

	return string(marshaledRes), nil
}

func SaveChatData(prop map[string]any) {

	_, err := db.SaveData(prop)

	fmt.Println(prop)

	if err != nil {
		log.Fatalf("Failed to save chat data %v", err)
	}
}
