package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type AddUserByChatIDUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewAddUserByChatIDUsecase(chatGateway gateway.ChatGateway) *AddUserByChatIDUsecase {
	return &AddUserByChatIDUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *AddUserByChatIDUsecase) Execute(ctx context.Context, input dto.AddUserByChatIDInputDTO) (*dto.AddUserByChatIDOutputDTO, error) {
	chat, err := u.ChatGateway.AddUserByChatID(ctx, input.ChatID, input.UserID)
	if err != nil {
		return nil, errors.New("failed to add user to the chat " + input.ChatID + ": " + err.Error())
	}

	return &dto.AddUserByChatIDOutputDTO{
		ID:        chat.ID.String(),
		ToUsers:   chat.ToUsers,
		FromUser:  chat.FromUser,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}, nil
}
