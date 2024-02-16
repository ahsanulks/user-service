package driver

import (
	"context"

	"userservice/internal/user/entity"
)

type UserGetterUsecase interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}
