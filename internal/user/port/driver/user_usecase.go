package driver

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, params *usecase.CreateUserParam) (id string, err error)
}
