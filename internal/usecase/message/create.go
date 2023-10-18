package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type CreateMessageUsecase struct {
	MessageGateway gateway.MessageGateway
}

func NewCreateMessageUsecase(messageGateway gateway.MessageGateway) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageGateway: messageGateway,
	}
}

func (mu *CreateMessageUsecase) Execute(ctx context.Context, message dto.CreateMessageInputDTO) (*dto.CreateMessageOutputDTO, error) {
	newMessage := entity.NewMessage(message.ChatID, message.Content, message.File)
	err := mu.MessageGateway.Create(ctx, newMessage)
	if err != nil {
		return nil, errors.New("failed to create message: " + err.Error())
	}
	return &dto.CreateMessageOutputDTO{
		ChatID:  newMessage.ChatId,
		Content: newMessage.Content,
		File:    newMessage.File,
		Created: newMessage.Created,
	}, nil
}
