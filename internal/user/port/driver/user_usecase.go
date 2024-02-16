package driver

import (
	"context"

	"userservice/internal/user/entity"
	"userservice/internal/user/param/request"
	"userservice/internal/user/param/response"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, params *request.CreateUser) (id string, err error)
	UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error)
	GenerateUserToken(ctx context.Context, params *request.GenerateUserTokenRequest) (*response.Token, error)
}
