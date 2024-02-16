package fake

import (
	"context"
	"database/sql"
	"errors"
	"time"

	customerror "userservice/internal/customError"
	"userservice/internal/user/entity"
	"userservice/internal/user/param/request"
	"userservice/internal/user/param/response"
	"userservice/internal/user/port/driver"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var (
	_ driver.UserUsecase       = new(FakeUserUsecase)
	_ driver.UserGetterUsecase = new(FakeUserUsecase)
)

type FakeUserUsecase struct {
	data     map[string]string
	dataById map[string]*entity.User
}

func NewFakeUserUsecase() *FakeUserUsecase {
	return &FakeUserUsecase{
		data:     make(map[string]string),
		dataById: make(map[string]*entity.User),
	}
}

// CreateUser implements driver.UserUsecase.
func (fu *FakeUserUsecase) CreateUser(ctx context.Context, params *request.CreateUser) (id string, err error) {
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
	fu.dataById[id] = &entity.User{
		ID:          id,
		FullName:    params.FullName,
		PhoneNumber: params.PhoneNumber,
		Password:    params.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return
}

func (fu FakeUserUsecase) GetIDByPhone(phone string) (id string, err error) {
	if id, ok := fu.data[phone]; ok {
		return id, nil
	}
	return "", errors.New("not Found")
}

// GetUserByID implements driver.UserGetterUsecase.
func (fu *FakeUserUsecase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	if id == "1232131" {
		return nil, sql.ErrNoRows
	}
	if data, ok := fu.dataById[id]; ok {
		return data, nil
	}
	return nil, errors.New("not Found")
}

// UpdateProfileByID implements driver.UserUsecase.
func (fu *FakeUserUsecase) UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error) {
	if id == "1232131" {
		return nil, errors.New("error")
	} else if id == "3333" {
		return nil, &pq.Error{
			Code: "23505",
		}
	}
	if data, ok := fu.dataById[id]; ok {
		return data, nil
	}
	return nil, errors.New("not Found")
}

// GenerateUserToken implements driver.UserUsecase.
func (*FakeUserUsecase) GenerateUserToken(ctx context.Context, params *request.GenerateUserTokenRequest) (*response.Token, error) {
	if params.PhoneNumber == "0000000" {
		return nil, customerror.NewValidationErrorWithMessage("authentication", "error")
	} else if params.PhoneNumber == "11111111" {
		return nil, errors.New("error")
	}
	return &response.Token{
		Token:     faker.Jwt(),
		ExpiresIn: 3600,
		Type:      "Bearer",
	}, nil
}
