package entity

type Chat struct {
	BasicEntity
	FromUser string
	ToUsers  []string
}

func NewChat(fromUser string, toUsers []string) *Chat {
	return &Chat{
		BasicEntity: NewBasicEntity(),
		FromUser:    fromUser,
		ToUsers:     toUsers,
	}
}
