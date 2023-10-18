package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type DeleteChatByIDUsecase struct {
	ChatGateway   gateway.ChatGateway
	MessageGateay gateway.MessageGateway
}

func NewDeleteChatByIDUsecase(chatGateway gateway.ChatGateway, messageGateway gateway.MessageGateway) *DeleteChatByIDUsecase {
	return &DeleteChatByIDUsecase{
		ChatGateway:   chatGateway,
		MessageGateay: messageGateway,
	}
}

func (u *DeleteChatByIDUsecase) Execute(ctx context.Context, input dto.DeleteChatByIDInputDTO) error {
	err := u.MessageGateay.DeleteAllMessageByChatID(ctx, input.ChatID)
	if err != nil {
		return errors.New("failed to delete messages: " + err.Error())
	}
	err = u.ChatGateway.DeleteChatByID(ctx, input.ChatID)
	if err != nil {
		return errors.New("failed to delete chat: " + err.Error())
	}

	return nil
}
