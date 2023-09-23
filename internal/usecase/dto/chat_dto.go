package dto

import (
	"time"
)

type MessageDTO struct {
	ChatId  string    `json:"chat_id"`
	Content string    `json:"content"`
	Files   [][]byte  `json:"files"`
	Created time.Time `json:"created"`
}

type ChatDTO struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string
	ToUsers   []string
	Messages  []MessageDTO
}

// Entrada de GetByChatId
type GetByChatIdInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Saída de GetByChatId
type GetByChatIdOutputDTO struct {
	Chat *ChatDTO `json:"chat"`
}

// Entrada de GetByUserId
type GetByUserIdInputDTO struct {
	UserID string `json:"user_id"`
}

// Saída de GetByUserId
type GetByUserIdOutputDTO struct {
	Chat *ChatDTO `json:"chat"`
}

// Entrada de GetToUsers
type GetToUsersInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Saída de GetToUsers
type GetToUsersOutputDTO struct {
	UserIDs []string `json:"user_ids"`
}

// Entrada de GetFromUser
type GetFromUserInputDTO struct {
	UserID string `json:"user_id"`
}

// Saída de GetFromUser
type GetFromUserOutputDTO struct {
	User *UserDTO `json:"user"`
}

// Entrada de CreateChat
type CreateChatInputDTO struct {
	Chat *ChatDTO `json:"chat"`
}

// Entrada de DeleteChatById
type DeleteChatByIdInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Entrada de AddMessage
type AddMessageInputDTO struct {
	Message *MessageDTO `json:"message"`
}

// Entrada de AddUser
type AddUserInputDTO struct {
	User *UserDTO `json:"user"`
}
