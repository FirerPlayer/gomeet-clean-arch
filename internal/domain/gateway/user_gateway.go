package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	// "github.com/firerplayer/gomeet-clean/internal/domain/entity"
)

type UserGateway interface {
	Create(ctx context.Context, user *entity.User) (string, error)
	DeleteUserByID(ctx context.Context, id string) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	ListAll(ctx context.Context, limit int) ([]*entity.User, error)
	UpdateUserByID(ctx context.Context, id string, user *entity.User) error
}
