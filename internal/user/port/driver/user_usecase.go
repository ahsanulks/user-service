package driver

import (
	"context"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/SawitProRecruitment/UserService/internal/user/param/response"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, params *request.CreateUser) (id string, err error)
	UpdateProfileByID(ctx context.Context, id string, params *request.UpdateProfile) (*entity.User, error)
	GenerateUserToken(ctx context.Context, params *request.GenerateUserTokenRequest) (*response.Token, error)
}
