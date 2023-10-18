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

func NewChatRepository(db *arangodb.DB, collectionName string) *ChatRepository {
	coll := db.FromCollection(collectionName)
	return &ChatRepository{
		arangodb:       db,
		collectionName: collectionName,
		dbCollection:   coll,
	}
}

// var byID = ""

func (cr *ChatRepository) Create(ctx context.Context, chat *entity.Chat) (string, error) {
	_, err := cr.dbCollection.InsertDocument(chat)
	if err != nil {
		return "", errors.New("failed to create chat: " + err.Error())
	}
	return chat.ID.String(), nil
}

const getChatById = "for chat in Chat filter chat.id == @id return chat"

func (cr *ChatRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	bindVars := map[string]interface{}{
		"id": chatID,
	}
	cursor, err := cr.arangodb.Database.Query(ctx, getChatById, bindVars)
	if err != nil {
		return nil, errors.New("failed to get chat by id: " + err.Error())
	}
	defer cursor.Close()
	var chat *entity.Chat
	_, err = cursor.ReadDocument(ctx, &chat)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
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

const deleteChatByID = "for chat in Chat filter chat.id == @chatID remove chat in Chat"

func (cr *ChatRepository) DeleteChatByID(ctx context.Context, chatID string) error {
	bindVars := map[string]interface{}{
		"chatID": chatID,
	}
	err := cr.dbCollection.ExecQuery(deleteChatByID, bindVars)
	if err != nil {
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
