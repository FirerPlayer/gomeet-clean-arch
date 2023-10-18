package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type DeleteByIDUsecase struct {
	UserGateway gateway.UserGateway
}

func NewDeleteByIDUsecase(userGateway gateway.UserGateway) *DeleteByIDUsecase {
	return &DeleteByIDUsecase{
		UserGateway: userGateway,
	}
}

func (u *DeleteByIDUsecase) Execute(ctx context.Context, input dto.DeleteUserByIDInputDTO) error {
	err := u.UserGateway.DeleteUserByID(ctx, input.ID)
	if err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}

	return nil
}
