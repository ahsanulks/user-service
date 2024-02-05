package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/test/fake"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_GenerateUserToken_CannotBindBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/token",
		strings.NewReader(`{"PhoneNumber": "123456789", "Password": "password123",}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.GenerateUserToken(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusBadRequest, rec.Code)
	assert.Equal("RequestBodyError", response.Type)
}

func TestServer_GenerateUserToken_WrongUsernamePassword(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/token",
		strings.NewReader(`{"PhoneNumber": "0000000", "Password": "password123"}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.GenerateUserToken(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusBadRequest, rec.Code)
	assert.Equal("ValidationError", response.Type)
}

func TestServer_GenerateUserToken_UnexpectedError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/token",
		strings.NewReader(`{"PhoneNumber": "11111111", "Password": "password123"}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.GenerateUserToken(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusInternalServerError, rec.Code)
	assert.Equal("InternalServerError", response.Type)
}

func TestServer_GenerateUserToken_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/users/token", strings.NewReader(`{"PhoneNumber": "123456789", "Password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.GenerateUserToken(ctx)
	assert.NoError(err)

	var response generated.CreateTokenResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)
	assert.Equal(http.StatusOK, rec.Code)
	assert.NotEmpty(response.AccessToken)
	assert.Equal("Bearer", response.Type)
}
