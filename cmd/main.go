package main

import (
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	server := newServer()
	e.Use(middleware.WithJwtAuth(server.TokenProvider))

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":8080"))
}

func newServer() *handler.Server {
	return handler.NewServer(&handler.ServerOptions{
		Conf: config.NewConfig(),
	})
}
