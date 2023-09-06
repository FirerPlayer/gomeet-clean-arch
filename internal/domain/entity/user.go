package entity

type User struct {
	BasicEntity
	Avatar []byte
	Name   string
	Email  string
	Bio    string
}

func NewUser(name, email, bio string) *User {
	return &User{
		BasicEntity: NewBasicEntity(),
		Avatar:      nil,
		Name:        name,
		Email:       email,
		Bio:         bio,
	}
}
