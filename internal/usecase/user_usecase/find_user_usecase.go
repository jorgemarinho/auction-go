package user_usecase

import (
	"context"

	"github.com/jorgemarinho/auction-go/internal/entity/user_entity"
	"github.com/jorgemarinho/auction-go/internal/internal_error"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
	FindUserByID(
		ctx context.Context,
		id string) (*UserOutputDTO, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserByID(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
