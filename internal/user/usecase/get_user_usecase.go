package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
)

type UserGetterUsecase struct {
	userGetter driven.UserGetter
}

func NewUserGetterUsecase(userGetter driven.UserGetter) *UserGetterUsecase {
	return &UserGetterUsecase{
		userGetter: userGetter,
	}
}

func (ug UserGetterUsecase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}
