package constants

import "fmt"

type Class int

const (
	ClassChat Class = iota
	ClassDocs
	ClassContent
)

func (c Class) String() string {
	switch c {
	case ClassChat:
		return "Chat"
	case ClassDocs:
		return "Docs"
	case ClassContent:
		return "Content"
	default:
		return fmt.Sprintf("Status(%d)", c)
	}
}

const DbClassChatDesc = "A text documument which is collection of chats between users and AI responses"
