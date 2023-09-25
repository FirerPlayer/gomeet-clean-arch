package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetToUsersUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewGetToUsersUsecase(chatGateway gateway.ChatGateway) *GetToUsersUsecase {
	return &GetToUsersUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *GetToUsersUsecase) Execute(ctx context.Context, input dto.GetToUsersInputDTO) ([]*dto.GetToUsersOutputDTO, error) {
	userIds, err := u.ChatGateway.GetToUsersByChatID(ctx, input.ChatID)
	if err != nil {
		return nil, errors.New("Failed to get to users: " + err.Error())
	}

	var out []*dto.GetToUsersOutputDTO
	for _, chat := range userIds {
		out = append(out, &dto.GetToUsersOutputDTO{
			UserID: chat,
		})
	}

	return out, nil

}
