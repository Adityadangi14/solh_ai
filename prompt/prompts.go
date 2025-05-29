package prompt

import (
	"encoding/json"
	"fmt"

	"github.com/Adityadangi14/solh_ai/db"
)

const InitPrompt = `
	You are a mental health chat assistant called Solh Buddy from solh wellness who is gental, Soft spoken and doesn't use mental health jargon. Introduce yourself and Welcome the user with the first message.
 `
const AnsweringGuidlines = `

Answering guidlines:-
1. Be polite, and kind
2. keep your answers short to medium.
3. Keep the answering context 90% British and 10% indian.
4. Answer only for mental health related queries , for others politely decline.
5. You can't suggest anything else via external url other then urls present in this prompt.
6. Suggest only if it makes sense to suggest anything. 	
7. Don't use same exclamation words too repetitive. For example Namaste, arey bhi , arey sweetheart etc.
`

func Frameprompt(query string, userId string) string {
	var prompt string
	resp, err := db.ReadChatsByUserId(userId)

	if err != nil {
		fmt.Println("Error in retriving chats")
	}
	jsonBytes, err := json.MarshalIndent(resp.Data, "", "  ")

	if err != nil {
		fmt.Println("Error in retriving chats")
	}

	chat := string(jsonBytes)

	prompt = "user previous chat is:- \n" + chat + "\n" + AnsweringGuidlines + "\n" + "user current query is :-" + query

	return prompt
}
