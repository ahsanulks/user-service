package fake

import (
	"errors"

	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/response"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
)

var _ driven.TokenProvider[*entity.User] = new(FakeTokenProvider)

type FakeTokenProvider struct{}

// Generate implements driven.TokenProvider.
func (*FakeTokenProvider) Generate(user *entity.User) (*response.Token, error) {
	if user.PhoneNumber == "======111" {
		return nil, errors.New("invalid")
	}
	return &response.Token{
		Token:     "1231313213213131",
		ExpiresIn: 3600,
		Type:      "Bearer",
	}, nil
}

func (*FakeTokenProvider) ValidateJWT(tokenString string) (map[string]interface{}, error) {
	return nil, nil
}
