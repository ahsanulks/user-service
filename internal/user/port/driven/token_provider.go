package driven

import (
	"github.com/SawitProRecruitment/UserService/internal/user/param/response"
)

type TokenProvider[T any] interface {
	Generate(data T) (*response.Token, error)
}
