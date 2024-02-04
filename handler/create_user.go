package handler

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/labstack/echo/v4"
)

// CreateUser implements generated.ServerInterface.
func (s *Server) CreateUser(ctx echo.Context) error {
	var params generated.CreateUserRequest
	if err := ctx.Bind(&params); err != nil {
		return parseError(ctx, err)
	}
	_, err := s.uu.CreateUser(ctx.Request().Context(), &usecase.CreateUserParam{
		PhoneNumber: params.PhoneNumber,
		FullName:    params.FullName,
		Password:    params.Password,
	})

	if err != nil {
		return parseError(ctx, err)
	}

	return nil
}
