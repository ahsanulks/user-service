package fake

import (
	"errors"

	"userservice/internal/user/entity"
	"userservice/internal/user/param/response"
	"userservice/internal/user/port/driven"
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
