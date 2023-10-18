package entity

type User struct {
	BasicEntity
	Avatar []byte `json:"avatar"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
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
