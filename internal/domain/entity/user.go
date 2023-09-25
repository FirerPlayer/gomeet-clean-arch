package entity

type User struct {
	BasicEntity
	Avatar []byte
	Name   string
	Email  string
	Bio    string
}

func NewUser(name, email, bio string, avatar []byte) *User {
	return &User{
		BasicEntity: NewBasicEntity(),
		Avatar:      avatar,
		Name:        name,
		Email:       email,
		Bio:         bio,
	}
}
