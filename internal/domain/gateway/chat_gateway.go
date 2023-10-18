package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
)

type ChatGateway interface {
	GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
	ListChatByUserID(ctx context.Context, userID string, limit int) ([]*entity.Chat, error)
	Create(ctx context.Context, chat *entity.Chat) (string, error)
	DeleteChatByID(ctx context.Context, chatID string) error
	AddUserByChatID(ctx context.Context, chatID string, userId string) (*entity.Chat, error)
}
