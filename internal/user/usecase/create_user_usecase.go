package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
)

const (
	costEncryption = 10
)

func (uu UserUsecase) CreateUser(ctx context.Context, params *request.CreateUser) (id string, err error) {
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

func (uu UserUsecase) UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error) {
	user := &entity.User{
		ID: id,
	}
	err := user.UpdateProfile(params.FullName, params.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return uu.userWriter.UpdateProfileByID(ctx, id, params)
}
