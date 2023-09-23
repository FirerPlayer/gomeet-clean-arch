package entity

import "time"

type Message struct {
	ChatId  string
	Content string
	Files   [][]byte
	Created time.Time
}

func NewMessage(chatId string, content string, files [][]byte) *Message {
	return &Message{
		ChatId:  chatId,
		Content: content,
		Files:   files,
		Created: time.Now(),
	}
}
