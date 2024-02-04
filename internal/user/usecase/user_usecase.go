package usecase

import "github.com/SawitProRecruitment/UserService/internal/user/port/driven"

type UserUsecase struct {
	userWriter driven.UserWriter
}

func NewUserUsecase(userWriter driven.UserWriter) *UserUsecase {
	return &UserUsecase{
		userWriter: userWriter,
	}
}
