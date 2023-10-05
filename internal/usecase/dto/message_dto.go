package dto

import "time"

type CreateMessageInputDTO struct {
	ChatID  string `json:"chat_id"`
	Content string `json:"content"`
	Files   []byte `json:"files"`
}
type CreateMessageOutputDTO struct {
	ChatID  string    `json:"chat_id"`
	Content string    `json:"content"`
	Files   []byte    `json:"files"`
	Created time.Time `json:"created"`
}

type ListMessageByChatIDInputDTO struct {
	ChatID string `json:"chat_id"`
	Limit  int    `json:"limit"`
}

type ListMessageByChatIDOutputDTO struct {
	ChatID  string    `json:"chat_id"`
	Content string    `json:"content"`
	Files   []byte    `json:"files"`
	Created time.Time `json:"created"`
}

type ListMessageByChatIDAndSearchInputDTO struct {
	ChatID string `json:"chat_id"`
	Query  string `json:"query"`
}

type ListMessageByChatIDAndSearchOutputDTO struct {
	ChatID  string    `json:"chat_id"`
	Content string    `json:"content"`
	Files   []byte    `json:"files"`
	Created time.Time `json:"created"`
}
