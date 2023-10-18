package repository

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
)

type MessageRepository struct {
	arangodb       *arangodb.DB
	collectionName string
	dbCollection   *arangodb.DBCollection
}

func NewMessageRepository(db *arangodb.DB, collectionName string) *MessageRepository {
	coll := db.FromCollection(collectionName)
	return &MessageRepository{
		arangodb:       db,
		collectionName: collectionName,
		dbCollection:   coll,
	}
}

// Create saves a message to the database.
//
// It takes a context and a message as parameters.
// It returns an error if the message fails to be created.
func (mr *MessageRepository) Create(ctx context.Context, message *entity.Message) error {
	_, err := mr.dbCollection.InsertDocument(message)
	if err != nil {
		return errors.New("Failed to create message: " + err.Error())
	}
	return nil
}

const listMessageByChatID = "for m in Message filter m.chatId == @chatID limit @limit return m"

// ListMessageByChatID retrieves a list of messages by chat ID.
//
// ctx: the context.Context object for managing the lifetime of the operation.
// chatID: the chat ID for which the messages should be retrieved.
// limit: the maximum number of messages to retrieve.
// []*entity.Message: a slice of entity.Message pointers representing the retrieved messages.
// error: an error if the operation fails.
func (mr *MessageRepository) ListMessageByChatID(ctx context.Context, chatID string, limit int) ([]*entity.Message, error) {
	bindVars := map[string]interface{}{
		"chatID": chatID,
		"limit":  limit,
	}
	cursor, err := mr.arangodb.Database.Query(ctx, listMessageByChatID, bindVars)
	if err != nil {
		return nil, errors.New("Failed to list messages: " + err.Error())
	}
	defer cursor.Close()
	var messages []*entity.Message
	for cursor.HasMore() {
		var message entity.Message
		_, err = cursor.ReadDocument(ctx, &message)
		if err != nil {
			return nil, errors.New("Failed to list messages: " + err.Error())
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

// func (mr *MessageRepository) ListMessageByChatIDAndSearch(ctx context.Context, chatID, search string) ([]*entity.Message, error)

const deleteAllMessageByChatID = "FOR m IN Message FILTER m.chatId == @chatID REMOVE m IN Message"

// DeleteAllMessageByChatID deletes all messages in the message repository
// that belong to a specific chat ID.
//
// ctx: the context object for the function
// chatID: the ID of the chat
// error: an error if the deletion fails
// Returns: an error if the deletion fails
func (mr *MessageRepository) DeleteAllMessageByChatID(ctx context.Context, chatID string) error {
	bindVars := map[string]interface{}{
		"chatID": chatID,
	}
	err := mr.dbCollection.ExecQuery(deleteAllMessageByChatID, bindVars)
	if err != nil {
		return errors.New("Failed to delete messages: " + err.Error())
	}
	return nil

}
