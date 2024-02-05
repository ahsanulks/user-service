package driven

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
)

type UserWriter interface {
	Create(ctx context.Context, user *entity.User) (id string, err error)
	UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error)
}
