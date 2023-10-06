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

const getChatById = "for chat in Chat filter chat.id == @id return chat"

func (cr *ChatRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	// bindVars := map[string]interface{}{
	// 	"id": chatID,
	// }
	// err := cr.collectionInstance.SelectQuery(getChatById)
	// // _, err := cr.collectionInstance.GetByKey(chatID, chat)
	// if err != nil {
	// 	return nil, errors.New("Failed to get chat by id: " + err.Error())
	// }
	// return &chats[0], nil
	return nil, nil
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
