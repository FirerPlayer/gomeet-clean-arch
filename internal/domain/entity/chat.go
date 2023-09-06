package entity

import "github.com/google/uuid"

type Chat struct {
	BasicEntity
	FromUser uuid.UUID    `json:"fromUser"`
	ToUsers  []*uuid.UUID `json:"toUsers"`
	Messages []Message    `json:"messages"`
	// Closed   bool
}

func NewChat(fromUser uuid.UUID, toUsers []*uuid.UUID) *Chat {
	return &Chat{
		BasicEntity: NewBasicEntity(),
		FromUser:    fromUser,
		ToUsers:     toUsers,
	}
}
