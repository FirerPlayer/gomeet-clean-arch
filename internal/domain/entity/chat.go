package entity

type Chat struct {
	BasicEntity
	FromUser string   `json:"fromUser"`
	ToUsers  []string `json:"toUsers"`
}

func NewChat(fromUser string, toUsers []string) *Chat {
	return &Chat{
		BasicEntity: NewBasicEntity(),
		FromUser:    fromUser,
		ToUsers:     toUsers,
	}
}
