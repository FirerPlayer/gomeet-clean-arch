package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetByIDUsecase struct {
	UserGateway gateway.UserGateway
}

func NewGetByIDUsecase(userGateway gateway.UserGateway) *GetByIDUsecase {
	return &GetByIDUsecase{
		UserGateway: userGateway,
	}
}

func (u *GetByIDUsecase) Execute(ctx context.Context, input dto.GetUserByIDInputDTO) (*dto.GetUserByIDOutputDTO, error) {
	user, err := u.UserGateway.GetUserByID(ctx, input.ID)
	if err != nil {
		return nil, errors.New("Failed to get user: " + err.Error())
	}

	return &dto.GetUserByIDOutputDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}
