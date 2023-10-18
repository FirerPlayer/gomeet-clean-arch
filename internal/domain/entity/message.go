package entity

import "time"

type Message struct {
	ChatId  string    `json:"chatId"`
	Content string    `json:"content"`
	File    []byte    `json:"file"`
	Created time.Time `json:"created"`
}

func NewMessage(chatId string, content string, file []byte) *Message {
	return &Message{
		ChatId:  chatId,
		Content: content,
		File:    file,
		Created: time.Now(),
	}
}
