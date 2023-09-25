package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetFromUserUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewGetFromUserUsecase(chatGateway gateway.ChatGateway) *GetFromUserUsecase {
	return &GetFromUserUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *GetFromUserUsecase) Execute(ctx context.Context, input dto.GetFromUserInputDTO) (*dto.GetFromUserOutputDTO, error) {
	user, err := u.ChatGateway.GetFromUserByChatID(ctx, input.UserID)
	if err != nil {
		return nil, errors.New("Failed to get from user: " + err.Error())
	}

	return &dto.GetFromUserOutputDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}
