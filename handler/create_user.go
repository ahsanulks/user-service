package handler

import (
	"net/http"

	"userservice/generated"
	"userservice/internal/user/param/request"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateUser implements generated.ServerInterface.
func (s *Server) CreateUser(ctx echo.Context) error {
	var params generated.CreateUserRequest
	if err := ctx.Bind(&params); err != nil {
		return parseError(ctx, err)
	}
	id, err := s.userUsecase.CreateUser(ctx.Request().Context(), &request.CreateUser{
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
