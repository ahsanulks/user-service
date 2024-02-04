package driven

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
)

// GetByID implements driven.UserGetter.
func (udb *UserDB) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}
