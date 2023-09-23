package dto

import "time"

type UserDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Entrada de CreateUser
type CreateUserInputDTO struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
	Avatar []byte `json:"avatar"`
}

// Saída de CreateUser
type CreateUserOutputDTO struct {
	ID string `json:"id"`
}

// Entrada de DeleteUserById
type DeleteUserByIdInputDTO struct {
	ID string `json:"id"`
}

// Entrada de GetUserById
type GetUserByIdInputDTO struct {
	ID string `json:"id"`
}

// Saída de GetUserById
type GetUserByIdOutputDTO struct {
	User *UserDTO `json:"user"`
}

// Saída de GetAllUsers
type GetAllUsersOutputDTO struct {
	Users []*UserDTO `json:"users"`
}

// Entrada de UpdateUserById
type UpdateUserByIdInputDTO struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
	Avatar []byte `json:"avatar"`
}
