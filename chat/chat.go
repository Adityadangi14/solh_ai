package chat

import (
	"context"
	"log"
	"time"

	"github.com/Adityadangi14/solh_ai/db"
	"github.com/Adityadangi14/solh_ai/initializers"
	"google.golang.org/genai"
)

type Chat struct {
	Query     string
	Answer    string
	UserID    string
	Timestamp time.Time
}

func (c *Chat) Map() map[string]any {
	return map[string]any{

		"query":     c.Query,
		"answer":    c.Answer,
		"userID":    c.UserID,
		"timestamp": c.Timestamp,
	}
}

// func StartChat() {
// 	ctx := context.Background()
// 	isInitPromptDone := false
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		if !isInitPromptDone {
// 			res, err := SendPrompt(ctx, prompt.InitPrompt)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			fmt.Println(res)

// 			isInitPromptDone = true

// 		} else {
// 			fmt.Print("You: ")
// 			userPrompt, _ := reader.ReadString('\n')
// 			userPrompt = strings.TrimSpace(userPrompt)

// 			userPrompt = prompt.Frameprompt(userPrompt)

// 			fmt.Println("===================================================")
// 			fmt.Println(userPrompt)
// 			fmt.Println("===================================================")

// 			if userPrompt == "" {
// 				continue
// 			}

// 			res, err := SendPrompt(ctx, userPrompt)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			fmt.Println("Sampada:", res)

// 			saveChatData(map[string]any{
// 				"userId":    1,
// 				"query":     userPrompt,
// 				"answer":    res,
// 				"timestamp": time.Now(),
// 			})
// 		}
// 	}
// }

func SendPrompt(ctx context.Context, prompt string, userId string) (string, error) {
	result, err := initializers.GemClient.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}
	res := result.Text()

	obj := Chat{Query: prompt, Answer: res, UserID: userId, Timestamp: time.Now()}
	saveChatData(obj.Map())
	return res, nil
}

func saveChatData(prop map[string]any) {

	_, err := db.SaveData(prop)

	if err != nil {
		log.Fatalln("Failed to save chat data")
	}
}
