package dto

import (
	"time"
)

type MessageDTO struct {
	ChatID  string    `json:"chatId"`
	Content string    `json:"content"`
	Files   [][]byte  `json:"files"`
	Created time.Time `json:"created"`
}

type ChatDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"fromUser"`
	ToUsers   []string  `json:"toUsers"`
}

// Entrada de GetByChatID
type GetByChatIDInputDTO struct {
	ChatID string `json:"chatId"`
}

// Saída de GetByChatID
type GetByChatIDOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"fromUser"`
	ToUsers   []string  `json:"toUsers"`
}

// Entrada de ListChatByUserID
type ListChatByUserIDInputDTO struct {
	UserID string `json:"userId"`
	Limit  int    `json:"limit"`
}

// Saída de ListChatByUserID
type ListChatByUserIDOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"fromUser"`
	ToUsers   []string  `json:"toUsers"`
}

// Entrada de GetToUsers
type GetToUsersInputDTO struct {
	ChatID string `json:"chatId"`
}

// Saída de GetToUsers
type GetToUsersOutputDTO struct {
	UserID string `json:"userId"`
}

// Entrada de GetFromUser
type GetFromUserInputDTO struct {
	UserID string `json:"userId"`
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
	FromUser string   `json:"fromUser"`
	ToUsers  []string `json:"toUsers"`
}

// Saida de CreateChat
type CreateChatOutputDTO struct {
	ChatID string `json:"chatId"`
}

// Entrada de DeleteChatById
type DeleteChatByIDInputDTO struct {
	ChatID string `json:"chatId"`
}

// Entrada de AddUser
type AddUserByChatIDInputDTO struct {
	ChatID string `json:"chatId"`
	UserID string `json:"userId"`
}

type AddUserByChatIDOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FromUser  string    `json:"fromUser"`
	ToUsers   []string  `json:"toUsers"`
}
