package driven

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/SawitProRecruitment/UserService/config"
	"github.com/SawitProRecruitment/UserService/internal/user/entity"
	"github.com/SawitProRecruitment/UserService/internal/user/param/response"
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrorInvalidToken = errors.New("token invalid")

var _ driven.TokenProvider[*entity.User] = new(UserTokenProvider)

type UserTokenProvider struct {
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	ExpiresSecond int
}

func NewUserTokenProvider(conf *config.JWT) *UserTokenProvider {
	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(conf.PrivateKey))
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(conf.PublicKey))
	if err != nil {
		panic(err)
	}
	return &UserTokenProvider{
		PrivateKey:    secretKey,
		PublicKey:     publicKey,
		ExpiresSecond: conf.ExpiresSecond,
	}
}

func (utp *UserTokenProvider) Generate(user *entity.User) (*response.Token, error) {
	jwtID, _ := uuid.NewRandom()
	claims := jwt.RegisteredClaims{
		Issuer:    "SawitPro",
		Subject:   user.ID,
		Audience:  []string{"user-service"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(utp.ExpiresSecond))),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        jwtID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(utp.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &response.Token{
		Token:     tokenString,
		ExpiresIn: int(time.Hour.Seconds()),
		Type:      "Bearer",
	}, nil
}

func (uc *UserTokenProvider) ValidateJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return uc.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrorInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorInvalidToken
	}

	return claims, nil
}
