package entity

import "github.com/google/uuid"

type Chat struct {
	BasicEntity
	FromUser uuid.UUID
	ToUsers  []*uuid.UUID
	Messages []Message
}

func NewChat(fromUser uuid.UUID, toUsers []*uuid.UUID) *Chat {
	return &Chat{
		BasicEntity: NewBasicEntity(),
		FromUser:    fromUser,
		ToUsers:     toUsers,
	}
}
