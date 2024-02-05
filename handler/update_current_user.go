package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/labstack/echo/v4"
)

func (s *Server) UpdateCurrentUserProfile(ctx echo.Context) error {
	var params generated.UpdateUserRequest
	if err := ctx.Bind(&params); err != nil {
		return parseError(ctx, err)
	}
	userID := ctx.Request().Context().Value("userID").(string)
	user, err := s.userUsecase.UpdateProfileByID(ctx.Request().Context(), userID, &request.UpdateProfile{
		PhoneNumber: params.PhoneNumber,
		FullName:    params.FullName,
	})

	if err != nil {
		return parseError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, generated.UpdateUserResponse{
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	})
}
