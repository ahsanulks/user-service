package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateUser implements generated.ServerInterface.
func (s *Server) CreateUser(ctx echo.Context) error {
	var params generated.CreateUserRequest
	if err := ctx.Bind(&params); err != nil {
		return parseError(ctx, err)
	}
	id, err := s.userUsecase.CreateUser(ctx.Request().Context(), &usecase.CreateUserParam{
		PhoneNumber: params.PhoneNumber,
		FullName:    params.FullName,
		Password:    params.Password,
	})

	if err != nil {
		return parseError(ctx, err)
	}

	uuid, _ := uuid.Parse(id)
	return ctx.JSON(http.StatusCreated, generated.CreateUserResponse{
		Id: uuid,
	})
}
