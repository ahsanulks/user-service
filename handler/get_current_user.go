package handler

import (
	"net/http"

	"userservice/generated"

	"github.com/labstack/echo/v4"
)

// GetCurrentUser implements generated.ServerInterface.
func (s *Server) GetCurrentUser(ctx echo.Context) error {
	userID := ctx.Request().Context().Value("userID").(string)
	user, err := s.userGetter.GetUserByID(ctx.Request().Context(), userID)
	if err != nil {
		return parseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, generated.GetUserResponse{
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	})
}
