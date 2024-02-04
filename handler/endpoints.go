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

// GetCurrentUser implements generated.ServerInterface.
func (*Server) GetCurrentUser(ctx echo.Context) error {
	panic("unimplemented")
}

// UpdateCurrentUserData implements generated.ServerInterface.
func (*Server) UpdateCurrentUserData(ctx echo.Context) error {
	panic("unimplemented")
}
