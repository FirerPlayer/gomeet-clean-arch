package repository

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
)

type ChatRepository struct {
	DB                 *arangodb.DB
	CollectionName     string
	collectionInstance *arangodb.DBCollection
}

func NewChatRepository(db *arangodb.DB, collectionName string) (*ChatRepository, error) {
	coll, err := db.FromCollection(collectionName)
	if err != nil {
		return nil, errors.New("error finding a collection: " + err.Error())
	}
	return &ChatRepository{
		DB:                 db,
		CollectionName:     collectionName,
		collectionInstance: coll,
	}, nil
}

// var byID = ""

func (r *ChatRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	// var chat entity.Chat
	// err := r.collectionInstance.SelectQuery()

	return nil, nil
}
