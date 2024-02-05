package usecase

import (
	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
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
