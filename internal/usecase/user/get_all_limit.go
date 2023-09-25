package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetAllLimitUsersUsecase struct {
	UserGateway gateway.UserGateway
}

func NewGetAllLimitUsersUsecase(userGateway gateway.UserGateway) *GetAllLimitUsersUsecase {
	return &GetAllLimitUsersUsecase{
		UserGateway: userGateway,
	}
}

func (u *GetAllLimitUsersUsecase) Execute(ctx context.Context, input dto.GetAllLimitUsersInputDTO) ([]*dto.GetAllLimitUsersOutputDTO, error) {
	result, err := u.UserGateway.ListAll(ctx, input.Limit)
	if err != nil {
		return nil, errors.New("Failed to get users: " + err.Error())
	}

	var out []*dto.GetAllLimitUsersOutputDTO
	for _, user := range result {
		out = append(out, &dto.GetAllLimitUsersOutputDTO{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			Bio:       user.Bio,
			Avatar:    user.Avatar,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return out, nil
}
