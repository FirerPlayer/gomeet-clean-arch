package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type DeleteAllMessagesByChatIDUsecase struct {
	MessageGateway gateway.MessageGateway
}

func NewDeleteAllMessagesByChatIDUsecase(messageGateway gateway.MessageGateway) *DeleteAllMessagesByChatIDUsecase {
	return &DeleteAllMessagesByChatIDUsecase{
		MessageGateway: messageGateway,
	}
}

func (u *DeleteAllMessagesByChatIDUsecase) Execute(ctx context.Context, input dto.DeleteAllMessageByChatIDInputDTO) error {
	err := u.MessageGateway.DeleteAllMessageByChatID(ctx, input.ChatID)
	if err != nil {
		return errors.New("failed to delete messages: " + err.Error())
	}
	return nil
}
