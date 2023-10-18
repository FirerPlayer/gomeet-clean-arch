package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type CreateChatUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewCreateChatUsecase(chatGateway gateway.ChatGateway) *CreateChatUsecase {
	return &CreateChatUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *CreateChatUsecase) Execute(ctx context.Context, input dto.CreateChatInputDTO) (*dto.CreateChatOutputDTO, error) {
	newChat := entity.NewChat(input.FromUser, input.ToUsers)
	chatId, err := u.ChatGateway.Create(ctx, newChat)
	if err != nil {
		return nil, errors.New("failed to create chat: " + err.Error())
	}

	return &dto.CreateChatOutputDTO{
		ChatID: chatId,
	}, nil

}
