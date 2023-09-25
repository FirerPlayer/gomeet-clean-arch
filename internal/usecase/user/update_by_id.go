package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type UpdateByIDUsecase struct {
	UserGateway gateway.UserGateway
}

func NewUpdateByIDUsecase(userGateway gateway.UserGateway) *UpdateByIDUsecase {
	return &UpdateByIDUsecase{
		UserGateway: userGateway,
	}
}

func (u *UpdateByIDUsecase) Execute(ctx context.Context, input dto.UpdateUserByIDInputDTO) error {
	err := u.UserGateway.UpdateUserByID(ctx, input.UserID,
		&entity.User{
			Name:   input.Name,
			Email:  input.Email,
			Bio:    input.Bio,
			Avatar: input.Avatar,
		},
	)
	if err != nil {
		return errors.New("Failed to update user: " + err.Error())
	}
	return nil
}
