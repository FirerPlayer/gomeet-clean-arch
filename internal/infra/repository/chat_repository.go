package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
)

type ChatRepository struct {
	DB                 *arangodb.DB
	CollectionName     string
	collectionInstance *arangodb.DBCollection
}

func NewChatRepository(db *arangodb.DB, collectionName string) (*ChatRepository, error) {
	coll := db.FromCollection(collectionName)
	return &ChatRepository{
		DB:                 db,
		CollectionName:     collectionName,
		collectionInstance: coll,
	}, nil
}

// var byID = ""

func (cr *ChatRepository) Create(ctx context.Context, chat *entity.Chat) (string, error) {
	meta, err := cr.collectionInstance.InsertDocument(chat)
	if err != nil {
		return "", errors.New("Failed to create chat: " + err.Error())
	}
	return meta.Key, nil
}

var getById = func(q string) string {
	return fmt.Sprintf(
		`for user in User 
			filter user.id == %s 
			return user`,
		q,
	)
}

func (cr *ChatRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	chats := make([]entity.Chat, 0)
	err := cr.collectionInstance.SelectQuery(getById(chatID), chats)
	// _, err := cr.collectionInstance.GetByKey(chatID, chat)
	if err != nil {
		return nil, errors.New("Failed to get chat by id: " + err.Error())
	}
	return &chats[0], nil
}

func (cr *ChatRepository) ListChatByUserID(ctx context.Context, userID string, limit int) ([]*entity.Chat, error) {
	return nil, nil
}
func (cr *ChatRepository) GetToUsersByChatID(ctx context.Context, chatID string) ([]string, error) {
	return nil, nil
}

func (cr *ChatRepository) GetFromUserByChatID(ctx context.Context, userID string) (*entity.User, error) {
	return nil, nil
}

func (cr *ChatRepository) DeleteChatByID(ctx context.Context, chatID string) error {
	return nil
}

func (cr *ChatRepository) AddUserByChatID(ctx context.Context, chatID string, userId string) error {
	return nil
}
