package repository

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
)

type ChatRepository struct {
	arangodb       *arangodb.DB
	collectionName string
	dbCollection   *arangodb.DBCollection
}

type ChatDocument struct {
	Key      string   `json:"_key"`
	Id       string   `json:"id"`
	FromUser string   `json:"fromUser"`
	ToUsers  []string `json:"toUsers"`
}

func NewChatDocument(chat *entity.Chat) *ChatDocument {
	return &ChatDocument{
		Key:      chat.ID.String(),
		Id:       chat.ID.String(),
		FromUser: chat.FromUser,
		ToUsers:  chat.ToUsers,
	}
}

func NewChatRepository(db *arangodb.DB, collectionName string) *ChatRepository {
	coll := db.FromCollection(collectionName)
	return &ChatRepository{
		arangodb:       db,
		collectionName: collectionName,
		dbCollection:   coll,
	}
}

func (cr *ChatRepository) Create(ctx context.Context, chat *entity.Chat) (string, error) {
	chatDocument := NewChatDocument(chat)
	_, err := cr.dbCollection.InsertDocument(chatDocument)
	if err != nil {
		return "", errors.New("failed to create chat: " + err.Error())
	}
	return chat.ID.String(), nil
}

func (cr *ChatRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	var chat *entity.Chat
	_, err := cr.dbCollection.GetByKey(chatID, &chat)
	if err != nil {
		if driver.IsDataSourceOrDocumentNotFound(err) {
			return nil, errors.New("chat not found with id: " + chatID)
		}
		return nil, errors.New("failed to get chat by id: " + err.Error())
	}
	return nil, nil
}

const listAllChatByUserID = "for chat in Chat filter chat.FromUser == @userID limit @limit return chat"

func (cr *ChatRepository) ListChatByUserID(ctx context.Context, userID string, limit int) ([]*entity.Chat, error) {
	bindVars := map[string]interface{}{
		"userID": userID,
		"limit":  limit,
	}
	cursor, err := cr.arangodb.Database.Query(ctx, listAllChatByUserID, bindVars)
	if err != nil {
		return nil, errors.New("failed to list all chats: " + err.Error())
	}
	defer cursor.Close()
	var chats []*entity.Chat
	for cursor.HasMore() {
		var chat *entity.Chat
		_, err := cursor.ReadDocument(ctx, &chat)
		if err != nil {
			return nil, errors.New("failed to retrieve all chats: " + err.Error())
		}
		chats = append(chats, chat)
	}
	if len(chats) == 0 {
		return nil, errors.New("no chats found with this user: " + userID)
	}
	return chats, nil
}

func (cr *ChatRepository) DeleteChatByID(ctx context.Context, chatID string) error {
	_, err := cr.dbCollection.DeleteDocument(chatID)
	if err != nil {
		if driver.IsDataSourceOrDocumentNotFound(err) {
			return errors.New("chat not found with id: " + chatID)
		}
		return errors.New("failed to delete chat: " + err.Error())
	}
	return nil
}

const addUserByChatID = `for c in Chat 
	for u in User 
		filter c.id == @chatID and u.id == @userID 
		update c with {toUsers: PUSH(c.toUsers, @userID)} in Chat 
		return NEW`

func (cr *ChatRepository) AddUserByChatID(ctx context.Context, chatID string, userId string) (*entity.Chat, error) {
	bindVars := map[string]interface{}{
		"chatID": chatID,
		"userID": userId,
	}
	cursor, err := cr.arangodb.Database.Query(ctx, addUserByChatID, bindVars)
	if err != nil {
		return nil, errors.New("failed to add user to chat: " + err.Error())
	}
	defer cursor.Close()
	var newChat *entity.Chat
	if cursor.HasMore() {
		_, err = cursor.ReadDocument(ctx, newChat)
		if err != nil {
			return nil, errors.New("failed to add user to chat: " + err.Error())
		}
		return newChat, nil
	}

	return nil, errors.New("chatID or userID not found")
}

// Colletion returns the collection instance of the UserRepository.
//
// No parameters.
// Returns *driver.Collection.
func (cr *ChatRepository) Colletion() *driver.Collection {
	return cr.dbCollection.Collection()
}

// Database returns the database instance of the UserRepository.
//
// No parameters.
// Returns a pointer to the driver.Database.
func (cr *ChatRepository) Database() *driver.Database {
	return &cr.arangodb.Database
}
