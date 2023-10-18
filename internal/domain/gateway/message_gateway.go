package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
)

type MessageGateway interface {
	Create(ctx context.Context, message *entity.Message) error
	ListMessageByChatID(ctx context.Context, chatID string, limit int) ([]*entity.Message, error)
	// ListMessageByChatIDAndSearch(ctx context.Context, chatID, search string) ([]*entity.Message, error)
	DeleteAllMessageByChatID(ctx context.Context, chatID string) error
}
