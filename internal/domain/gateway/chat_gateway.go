package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
)

type ChatGateway interface {
	GetByChatId(ctx context.Context, chatId string) (*entity.Chat, error)
	GetByUserId(ctx context.Context, userId string) (*entity.Chat, error)
	GetToUsers(ctx context.Context, chatId string) ([]string, error)
	GetFromUser(ctx context.Context, userId string) (*entity.User, error)
	Create(ctx context.Context, chat *entity.Chat) error
	DeleteById(ctx context.Context, chatId string) error
	AddMessage(ctx context.Context, message *entity.Message) error
	AddUser(ctx context.Context, user *entity.User) error
}
