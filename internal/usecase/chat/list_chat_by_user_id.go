package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type ListChatByUserIDUsecase struct {
	ChatGateway gateway.ChatGateway
}

func NewListChatByUserIDUsecase(chatGateway gateway.ChatGateway) *ListChatByUserIDUsecase {
	return &ListChatByUserIDUsecase{
		ChatGateway: chatGateway,
	}
}

func (u *ListChatByUserIDUsecase) Execute(ctx context.Context, input dto.ListChatByUserIDInputDTO) ([]*dto.ListChatByUserIDOutputDTO, error) {
	result, err := u.ChatGateway.ListChatByUserID(ctx, input.UserID, input.Limit)
	if err != nil {
		return nil, errors.New("Failed to get users: " + err.Error())
	}
	var out []*dto.ListChatByUserIDOutputDTO
	for _, chat := range result {
		out = append(out, &dto.ListChatByUserIDOutputDTO{
			ID:        chat.ID.String(),
			ToUsers:   chat.ToUsers,
			FromUser:  chat.FromUser,
			CreatedAt: chat.CreatedAt,
			UpdatedAt: chat.UpdatedAt,
		})
	}
	return out, nil

}
