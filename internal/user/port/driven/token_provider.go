package driven

import (
	"userservice/internal/user/param/response"
)

type TokenProvider[T any] interface {
	Generate(data T) (*response.Token, error)
	ValidateJWT(tokenString string) (map[string]interface{}, error)
}
