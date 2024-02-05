package driven

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
)

type UserGetter interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
}
