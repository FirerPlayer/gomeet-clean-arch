package dto

import (
	"time"
)

type MessageDTO struct {
	ChatID  string    `json:"chat_id"`
	Content string    `json:"content"`
	Files   [][]byte  `json:"files"`
	Created time.Time `json:"created"`
}

type ChatDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"from_user"`
	ToUsers   []string  `json:"to_users"`
}

// Entrada de GetByChatID
type GetByChatIDInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Saída de GetByChatID
type GetByChatIDOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"from_user"`
	ToUsers   []string  `json:"to_users"`
}

// Entrada de ListChatByUserID
type ListChatByUserIDInputDTO struct {
	UserID string `json:"user_id"`
	Limit  int    `json:"limit"`
}

// Saída de ListChatByUserID
type ListChatByUserIDOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"from_user"`
	ToUsers   []string  `json:"to_users"`
}

// Entrada de GetToUsers
type GetToUsersInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Saída de GetToUsers
type GetToUsersOutputDTO struct {
	UserID string `json:"user_id"`
}

// Entrada de GetFromUser
type GetFromUserInputDTO struct {
	UserID string `json:"user_id"`
}

// Saída de GetFromUser
type GetFromUserOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Entrada de CreateChat
type CreateChatInputDTO struct {
	FromUser string   `json:"from_user"`
	ToUsers  []string `json:"to_users"`
}

// Saida de CreateChat
type CreateChatOutputDTO struct {
	ChatID string `json:"chat_id"`
}

// Entrada de DeleteChatById
type DeleteChatByIDInputDTO struct {
	ChatID string `json:"chat_id"`
}

// Entrada de AddUser
type AddUserByChatIDInputDTO struct {
	ChatID string `json:"chat_id"`
	UserID string `json:"user_id"`
}
