package driver

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
)

type UserGetterUsecase interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}
