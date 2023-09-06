package entity

import "github.com/google/uuid"

type Message struct {
	ChatId  uuid.UUID
	Content string
	Files   [][]byte
}

func NewMessage(chatId uuid.UUID, content string, files [][]byte) *Message {
	return &Message{
		ChatId:  chatId,
		Content: content,
		Files:   files,
	}
}
