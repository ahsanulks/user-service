package main

import (
	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	custommiddleware "github.com/SawitProRecruitment/UserService/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	server := newServer()

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		custommiddleware.WithJwtAuth(server.TokenProvider),
	)

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":8080"))
}

func newServer() *handler.Server {
	return handler.NewServer(&handler.ServerOptions{
		Conf: config.NewConfig(),
	})
}
