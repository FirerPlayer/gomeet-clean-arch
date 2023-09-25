package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type ListMessageByChatIDUsecase struct {
	MessageGateway gateway.MessageGateway
}

func NewListMessageByChatIDUsecase(messageGateway gateway.MessageGateway) *ListMessageByChatIDUsecase {
	return &ListMessageByChatIDUsecase{
		MessageGateway: messageGateway,
	}
}

func (u *ListMessageByChatIDUsecase) Execute(ctx context.Context, input dto.ListMessageByChatIDInputDTO) ([]*dto.ListMessageByChatIDOutputDTO, error) {
	result, err := u.MessageGateway.ListMessageByChatID(ctx, input.ChatID, input.Limit)
	if err != nil {
		return nil, errors.New("Failed to get users: " + err.Error())
	}

	var out []*dto.ListMessageByChatIDOutputDTO
	for _, message := range result {
		out = append(out, &dto.ListMessageByChatIDOutputDTO{
			ChatID:  message.ChatId,
			Content: message.Content,
			Files:   message.Files,
			Created: message.Created,
		})
	}

	return out, nil

}
