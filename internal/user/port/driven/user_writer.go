package driven

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
)

type UserWriter interface {
	Create(ctx context.Context, user *entity.User) (id string, err error)
}
