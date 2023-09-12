package gateway

import (
	"context"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	// "github.com/firerplayer/gomeet-clean/internal/domain/entity"
)

type UserGateway interface {
	Create(ctx context.Context, user *entity.User) error
	DeleteById(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	UpdateById(ctx context.Context, id string, user *entity.User) error
}
