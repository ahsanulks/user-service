package handler

import (
	"github.com/SawitProRecruitment/UserService/adapter/driven"
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driver"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
)

type Server struct {
	uu driver.UserUsecase
}

type ServerOptions struct {
	Conf *config.ApplicationConfig
}

func NewServer(opt *ServerOptions) *Server {
	db := driven.NewPostgreConnection(&opt.Conf.Postgres)
	userDB := driven.NewUserDB(db)

	return &Server{
		uu: usecase.NewUserUsecase(userDB, new(driven.BcyrpEncryption)),
	}
}
