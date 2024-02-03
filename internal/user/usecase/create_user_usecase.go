package usecase

import "context"

type CreateUserParam struct {
	PhoneNumber string
	FullName    string
	Password    string
}

func (uu UserUsecase) CreateUser(ctx context.Context, param *CreateUserParam) (id int, err error) {
	return
}
