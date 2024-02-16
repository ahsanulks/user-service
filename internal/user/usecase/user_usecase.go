package usecase

import (
	"userservice/internal/user/entity"
	"userservice/internal/user/port/driven"
)

type UserUsecase struct {
	userWriter    driven.UserWriter
	encryptor     driven.Encyptor
	userGetter    driven.UserGetter
	tokenProvider driven.TokenProvider[*entity.User]
}

func NewUserUsecase(
	userWriter driven.UserWriter,
	encryptor driven.Encyptor,
	userGetter driven.UserGetter,
	tokenProvider driven.TokenProvider[*entity.User],
) *UserUsecase {
	return &UserUsecase{
		userWriter:    userWriter,
		encryptor:     encryptor,
		userGetter:    userGetter,
		tokenProvider: tokenProvider,
	}
}
