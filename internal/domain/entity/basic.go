package entity

import (
	"time"

	"github.com/google/uuid"
)

type BasicEntity struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewBasicEntity() BasicEntity {
	return BasicEntity{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
