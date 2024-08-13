package user_entity

import (
	"context"

	"github.com/jorgemarinho/auction-go/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserByID(
		ctx context.Context,
		userId string) (*User, *internal_error.InternalError)
}
