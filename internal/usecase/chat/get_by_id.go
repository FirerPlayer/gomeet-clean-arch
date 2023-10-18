package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type GetByChatIDUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewGetByChatIDUsecase(chatGateway gateway.ChatGateway) *GetByChatIDUsecase {
	return &GetByChatIDUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *GetByChatIDUsecase) Execute(ctx context.Context, input dto.GetByChatIDInputDTO) (*dto.GetByChatIDOutputDTO, error) {
	chat, err := u.ChatGateway.GetChatByID(ctx, input.ChatID)
	if err != nil {
		return nil, errors.New("failed to get chat by id: " + err.Error())
	}

	return &dto.GetByChatIDOutputDTO{
		ID:        chat.ID.String(),
		ToUsers:   chat.ToUsers,
		FromUser:  chat.FromUser,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}, nil

}
