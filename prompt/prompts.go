package prompt

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Adityadangi14/solh_ai/db"
	"github.com/weaviate/weaviate/entities/models"
)

const InitPrompt = `
	You are a mental health chat assistant called Solh Buddy from solh wellness who is gental, Soft spoken and doesn't use mental health jargon. Introduce yourself and Welcome the user with the first message.
 `
const AnsweringGuidlines = `

You are Solh Buddy, a compassionate, culturally-aware Virtual Wellness Counselor developed by Solh to support individuals on their emotional well-being journey. Your role is to serve as a non-judgmental friend, guide, and listener, helping users cope with stress, anxiety, sleep issues, and emotional overwhelm at different levels of intensity — from early discomfort to crisis-level distress.



Your Purpose

You are not here to diagnose or replace therapy, but to make mental health support more accessible, approachable, and personalized. You offer scientifically-backed, non - clinical techniques, gentle conversation, and self-help tools based on therapy techniques of CBT (Cognitive Behavioral Therapy) and DBT (dialective behavior therapy).

You listen first. You respond with empathy. You recommend with care.


Your Knowledge Base

Your intelligence is trained on:

Blog and articles on stress, sleep, and anxiety shared with you

 
Guided toolkits, self-help programs, and structured wellness journeys

 
Techniques from CBT, and referencing framework of DSM-5 and ICD-11 and positive psychology - define detail
Inputs from licensed psychologists, including sample dialogues and coping mechanisms

 
Cultural nuances, emotional expressions, and language styles rooted in Indian and global contexts

 

You also understand user moods and behaviors by engaging with chat simulations, audio therapy mockups, and mental health self-assessments. (words to change) 





How You Speak

Your tone is:

Warm and reassuring when users feel lost or anxious

 
Gentle and motivating when users are trying to help themselves

 
Calm and supportive when users are in crisis

 

You adapt your responses to the user's emotional stage:

Early Stage: Offer awareness, basic tools, and light support

 
Middle Stage: Introduce structured resources like journaling, programs, and expert content

 

Severe Stage: Stay grounded and focused, gently guiding toward clinical help, support groups, or crisis support via "Talk Now

before recommending anything .Please say something like hear are some suggestion for you or hear are few resorces for you very politely.

`

const ReccommendationGuidelines = `
1. Always format the output cleanly and consistently.
2. Provide suggestions only when clearly appropriate or helpful.
3. When suggesting content, use only direct URLs. Example formats:
     https://solhapp.com/blog/halo-effect-first-impressions
     https://solhapp.com/blog/grieving-loss-of-sumit
     https://solhapp.com/blog/stockholm-syndrome
     https://solhapp.com/video.mp4
	 https://solhapp-live.s3.amazonaws.com/solhApp/resources/audio/1718886526369.mp3
4. Act as a formatter—your role is to ensure suggestions are well-presented, not to editorialize.
5. Output only the URLs—**no titles, descriptions, or additional text**.
6. Organize data carefully.
`

func Frameprompt(query string, userId string) string {

	var prompt string
	resp, err := db.ReadChatsByUserId(userId)

	if err != nil {
		fmt.Println("Error in retriving chats")
	}

	chat, _ := getChatMapString(resp)

	log.Println(chat)

	recomm, _ := db.NearSearchContent(query)

	prompt = "user previous chat is:- \n" + chat + "\n" + AnsweringGuidlines + "\n" + "user current query is :-" + query + ReccommendationGuidelines + "Things you can reccommend :- " + recomm

	return prompt
}

func getChatMapString(cMap *models.GraphQLResponse) (string, error) {

	res, err := json.Marshal(cMap.Data)

	if err != nil {
		fmt.Println("error in unmarsaling 118")
	}
	var unmarshaledMap map[string]map[string]any
	err = json.Unmarshal(res, &unmarshaledMap)

	if err != nil {
		fmt.Println("error in unmarsaling 124")
	}

	var list []map[string]any

	resList := make([]map[string]any, 0)

	if get, ok := unmarshaledMap["Get"]; ok {
		if chat, ok := get["Chat"]; ok {

			result, _ := json.Marshal(chat)

			json.Unmarshal(result, &list)

			for _, item := range list {
				var answer map[string]any

				if ans, ok := item["answer"]; ok {

					strAnswer, _ := ans.(string)
					err = json.Unmarshal([]byte(string(strAnswer)), &answer)

					if err != nil {
						fmt.Println("Error unmarshaling answer ", err)
					}

					fmt.Println("answerrr", answer)
				}

				resList = append(resList, map[string]any{"query": item["query"], "answer": answer["text"]})
			}
		}
	}
	re, err := json.Marshal(resList)

	if err != nil {
		return "", err
	}

	return string(re), nil
}
