package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetByEmailUsecase struct {
	UserGateway gateway.UserGateway
}

func NewGetByEmailUsecase(userGateway gateway.UserGateway) *GetByEmailUsecase {
	return &GetByEmailUsecase{
		UserGateway: userGateway,
	}
}

func (u *GetByEmailUsecase) Execute(ctx context.Context, input dto.GetUserByEmailInputDTO) (*dto.GetUserByEmailOutputDTO, error) {
	user, err := u.UserGateway.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("Failed to get user by email: " + err.Error())
	}

	return &dto.GetUserByEmailOutputDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}
