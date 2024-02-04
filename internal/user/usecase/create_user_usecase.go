package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
)

type CreateUserParam struct {
	PhoneNumber string
	FullName    string
	Password    string
}

const (
	costEncryption = 10
)

func (uu UserUsecase) CreateUser(ctx context.Context, params *CreateUserParam) (id string, err error) {
	user, err := entity.NewUser(params.FullName, params.PhoneNumber, params.Password)
	if err != nil {
		return id, err
	}

	encryptedPassword, err := uu.encryptor.Encrypt([]byte(user.Password), costEncryption)
	if err != nil {
		return id, err
	}
	user.Password = string(encryptedPassword)
	return uu.userWriter.Create(ctx, user)
}
