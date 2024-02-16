package driven

import (
	"context"

	"userservice/internal/user/entity"
	"userservice/internal/user/param/request"
)

type UserWriter interface {
	Create(ctx context.Context, user *entity.User) (id string, err error)
	UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error)
	UpdateUserToken(ctx context.Context, userId string) error
}
