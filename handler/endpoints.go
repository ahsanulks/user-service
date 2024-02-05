package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

var _ generated.ServerInterface = new(Server)

// GenerateUserToken implements generated.ServerInterface.
func (*Server) GenerateUserToken(ctx echo.Context) error {
	panic("unimplemented")
}
