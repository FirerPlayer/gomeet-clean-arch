package dto

import "time"

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

// Entrada de DeleteUserByID
type DeleteUserByIDInputDTO struct {
	ID string `json:"id"`
}

// Entrada de GetUserByID
type GetUserByIDInputDTO struct {
	ID string `json:"id"`
}

// Saída de GetUserByID
type GetUserByIDOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Entrada de GetAllLimitUsers
type GetAllLimitUsersInputDTO struct {
	Limit int `json:"limit"`
}

// Saída de GetAllUsers
type GetAllLimitUsersOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Entrada de UpdateUserByID
type UpdateUserByIDInputDTO struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
	Avatar []byte `json:"avatar"`
}

type UpdateUserByIDOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetUserByEmailInputDTO struct {
	Email string `json:"email"`
}

type GetUserByEmailOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
