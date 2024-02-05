package handler

import (
	"github.com/SawitProRecruitment/UserService/adapter/driven"
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driver"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
)

type Server struct {
	TokenProvider *driven.UserTokenProvider

	userUsecase driver.UserUsecase
	userGetter  driver.UserGetterUsecase
}

type ServerOptions struct {
	Conf *config.ApplicationConfig
}

func NewServer(opt *ServerOptions) *Server {
	db := driven.NewPostgreConnection(&opt.Conf.Postgres)
	userDB := driven.NewUserDB(db)
	tokenProvider := driven.NewUserTokenProvider(&opt.Conf.JWT)

	return &Server{
		userUsecase: usecase.NewUserUsecase(
			userDB,
			new(driven.BcyrpEncryption),
			userDB,
			tokenProvider,
		),
		userGetter:    usecase.NewUserGetterUsecase(userDB),
		TokenProvider: tokenProvider,
	}
}
