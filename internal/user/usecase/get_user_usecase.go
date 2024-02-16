package usecase

import (
	"context"

	"userservice/internal/user/entity"
	"userservice/internal/user/port/driven"
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
	return ug.userGetter.GetByID(ctx, id)
}
