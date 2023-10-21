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

func (u *UpdateByIDUsecase) Execute(ctx context.Context, input dto.UpdateUserByIDInputDTO) (*dto.UpdateUserByIDOutputDTO, error) {
	updatedUser, err := u.UserGateway.UpdateUserByID(ctx, input.UserID,
		&entity.User{
			Name:   input.Name,
			Email:  input.Email,
			Bio:    input.Bio,
			Avatar: input.Avatar,
		},
	)
	if err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}
	return &dto.UpdateUserByIDOutputDTO{
		ID:        updatedUser.ID.String(),
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Bio:       updatedUser.Bio,
		Avatar:    updatedUser.Avatar,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}, nil
}
