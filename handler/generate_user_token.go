package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/user/param/request"
	"github.com/labstack/echo/v4"
)

var _ generated.ServerInterface = new(Server)

// GenerateUserToken implements generated.ServerInterface.
func (s *Server) GenerateUserToken(ctx echo.Context) error {
	var params generated.CreateTokenRequest
	if err := ctx.Bind(&params); err != nil {
		return parseError(ctx, err)
	}
	token, err := s.userUsecase.GenerateUserToken(ctx.Request().Context(), &request.GenerateUserTokenRequest{
		PhoneNumber: params.PhoneNumber,
		Password:    params.Password,
	})
	if err != nil {
		return parseError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, generated.CreateTokenResponse{
		AccessToken: token.Token,
		ExpiresIn:   int32(token.ExpiresIn),
		Type:        token.Type,
	})
}
