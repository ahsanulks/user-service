package usecase

import (
	"context"

	customerror "userservice/internal/custom_error"
	"userservice/internal/user/entity"
	"userservice/internal/user/param/request"
	"userservice/internal/user/param/response"
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

func (uu UserUsecase) GenerateUserToken(ctx context.Context, params *request.GenerateUserTokenRequest) (*response.Token, error) {
	user, err := uu.userGetter.GetByPhoneNumber(ctx, params.PhoneNumber)
	if err != nil {
		return nil, customerror.NewValidationErrorWithMessage("authentication", "wrong phone number/password")
	}

	err = uu.encryptor.CompareEncryptedAndData([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, customerror.NewValidationErrorWithMessage("authentication", "wrong phone number/password")
	}

	token, err := uu.tokenProvider.Generate(user)
	if err != nil {
		return nil, err
	}
	err = uu.userWriter.UpdateUserToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
