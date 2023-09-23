package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
)

type MessageGateway interface {
	Create(ctx context.Context, message *entity.Message) error
	GetAllByChatId(ctx context.Context, chatId string) ([]*entity.Message, error)
	GetAllByChatIdAndSearch(ctx context.Context, chatId, search string) ([]*entity.Message, error)
	DeleteAllByChatId(ctx context.Context, chatId string) error
}
