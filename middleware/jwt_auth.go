package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"github.com/labstack/echo/v4"
)

// because we cannot add middleware 1 by 1
// the path is generated from openapi
var (
	unprotectedPath = map[string][]string{
		"POST": {
			"/api/v1/users",
			"/api/v1/users/token",
		},
	}
)

func WithJwtAuth(tokenProvider driven.TokenProvider[*entity.User]) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if accessUnprotectedPath(c) {
				return next(c)
			}

			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Unauthorized"})
			}

			token := strings.TrimPrefix(tokenString, "Bearer ")
			claims, err := tokenProvider.ValidateJWT(token)
			if err != nil {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid token"})
			}

			userID, _ := claims["sub"].(string)
			newContext := context.WithValue(c.Request().Context(), "userID", userID)

			c.SetRequest(c.Request().WithContext(newContext))

			return next(c)
		}
	}
}

func accessUnprotectedPath(c echo.Context) bool {
	path := c.Path()
	method := c.Request().Method
	for _, unprotectedPath := range unprotectedPath[method] {
		if unprotectedPath == path {
			return true
		}
	}
	return false
}
