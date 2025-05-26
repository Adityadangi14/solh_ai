package prompt

import (
	"encoding/json"
	"fmt"

	"github.com/Adityadangi14/solh_ai/db"
)

const InitPrompt = `
	You are a mental health chat assistant called Sampada from solh wellness who is gental, Soft spoken and doesn't use mental health jargon. Introduce yourself and Welcome the user with the first message.
 `
const answeringGuidlines = `

Answering guidlines:-
1. Be polite, and kind
2. keep your answers short to medium.
3. Keep the answering context Indian.
4. Answer only for mental health related queries , for others politely decline.
5. You can't suggest anything else via external url other then urls present in this prompt.
6. Suggest only if it makes sense to suggest anything. 	

`

func Frameprompt(query string) string {
	var prompt string
	resp, err := db.GetPreviousChat()

	if err != nil {
		fmt.Println("Error in retriving chats")
	}
	jsonBytes, err := json.MarshalIndent(resp.Data, "", "  ")

	if err != nil {
		fmt.Println("Error in retriving chats")
	}

	chat := string(jsonBytes)

	prompt = "user previous chat is:- \n" + chat + "\n" + answeringGuidlines + "\n" + "user current query is :-" + query

	return prompt
}
