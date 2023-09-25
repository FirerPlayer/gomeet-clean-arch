package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/domain/gateway"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
)

type CreateUserUsecase struct {
	UserGateway gateway.UserGateway
}

func NewCreateUserUsecase(userGateway gateway.UserGateway) *CreateUserUsecase {
	return &CreateUserUsecase{
		UserGateway: userGateway,
	}
}

func (u *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserInputDTO) (*dto.CreateUserOutputDTO, error) {

	newUser := entity.NewUser(input.Name, input.Email, input.Bio, input.Avatar)
	id, err := u.UserGateway.Create(ctx, newUser)
	if err != nil {
		return nil, errors.New("Failed to create user: " + err.Error())
	}

	return &dto.CreateUserOutputDTO{
		ID: id,
	}, nil
}
