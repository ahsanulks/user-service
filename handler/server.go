package handler

import (
	"userservice/config"
	"userservice/infrastructure"
	"userservice/internal/user/port/driver"
	"userservice/internal/user/usecase"
)

type Server struct {
	TokenProvider *infrastructure.UserTokenProvider

	userUsecase driver.UserUsecase
	userGetter  driver.UserGetterUsecase
}

type ServerOptions struct {
	Conf *config.ApplicationConfig
}

func NewServer(opt *ServerOptions) *Server {
	db := infrastructure.NewPostgreConnection(&opt.Conf.Postgres)
	userDB := infrastructure.NewUserDB(db)
	tokenProvider := infrastructure.NewUserTokenProvider(&opt.Conf.JWT)

	return &Server{
		userUsecase: usecase.NewUserUsecase(
			userDB,
			new(infrastructure.BcyrpEncryption),
			userDB,
			tokenProvider,
		),
		userGetter:    usecase.NewUserGetterUsecase(userDB),
		TokenProvider: tokenProvider,
	}
}
