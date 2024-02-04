package fake

import (
	"context"
	"errors"

	customerror "github.com/SawitProRecruitment/UserService/internal/customError"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driver"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/google/uuid"
)

var _ driver.UserUsecase = new(FakeUserUsecase)

type FakeUserUsecase struct {
	data map[string]string
}

func NewFakeUserUsecase() *FakeUserUsecase {
	return &FakeUserUsecase{
		data: make(map[string]string),
	}
}

// CreateUser implements driver.UserUsecase.
func (fu *FakeUserUsecase) CreateUser(ctx context.Context, params *usecase.CreateUserParam) (id string, err error) {
	if len(params.PhoneNumber) == 0 {
		customErr := customerror.NewValidationError()
		customErr.AddError("phone number", "phone empty")
		return "", customErr
	}

	if params.PhoneNumber == "11111111" {
		return "", errors.New("unexpected")
	}

	id = uuid.New().String()
	fu.data[params.PhoneNumber] = id
	return
}

func (fu FakeUserUsecase) GetIDByPhone(phone string) (id string, err error) {
	if id, ok := fu.data[phone]; ok {
		return id, nil
	}
	return "", errors.New("not Found")
}
