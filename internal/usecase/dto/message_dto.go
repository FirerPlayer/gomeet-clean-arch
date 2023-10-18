package dto

import "time"

type CreateMessageInputDTO struct {
	ChatID  string `json:"chatId"`
	Content string `json:"content"`
	File    []byte `json:"file"`
}
type CreateMessageOutputDTO struct {
	ChatID  string    `json:"chatId"`
	Content string    `json:"content"`
	File    []byte    `json:"file"`
	Created time.Time `json:"created"`
}

type ListMessageByChatIDInputDTO struct {
	ChatID string `json:"chatId"`
	Limit  int    `json:"limit"`
}

type ListMessageByChatIDOutputDTO struct {
	ChatID  string    `json:"chatId"`
	Content string    `json:"content"`
	File    []byte    `json:"file"`
	Created time.Time `json:"created"`
}

type DeleteAllMessageByChatIDInputDTO struct {
	ChatID string `json:"chatId"`
}

// type ListMessageByChatIDAndSearchInputDTO struct {
// 	ChatID string `json:"chatId"`
// 	Query  string `json:"query"`
// }

// type ListMessageByChatIDAndSearchOutputDTO struct {
// 	ChatID  string    `json:"chatId"`
// 	Content string    `json:"content"`
// 	File    []byte    `json:"file"`
// 	Created time.Time `json:"created"`
// }
