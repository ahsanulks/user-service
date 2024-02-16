package handler

import (
	"userservice/config"
	"userservice/internal/user/port/driver"
	"userservice/internal/user/usecase"
	"userservice/repository"
)

type Server struct {
	TokenProvider *repository.UserTokenProvider

	userUsecase driver.UserUsecase
	userGetter  driver.UserGetterUsecase
}

type ServerOptions struct {
	Conf *config.ApplicationConfig
}

func NewServer(opt *ServerOptions) *Server {
	db := repository.NewPostgreConnection(&opt.Conf.Postgres)
	userDB := repository.NewUserDB(db)
	tokenProvider := repository.NewUserTokenProvider(&opt.Conf.JWT)

	return &Server{
		userUsecase: usecase.NewUserUsecase(
			userDB,
			new(repository.BcyrpEncryption),
			userDB,
			tokenProvider,
		),
		userGetter:    usecase.NewUserGetterUsecase(userDB),
		TokenProvider: tokenProvider,
	}
}
