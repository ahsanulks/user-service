package handler

import (
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driver"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/SawitProRecruitment/UserService/repository"
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
