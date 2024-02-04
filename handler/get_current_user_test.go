package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/user/usecase"
	"github.com/SawitProRecruitment/UserService/test/fake"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUser_Success(t *testing.T) {
	fu := fake.NewFakeUserUsecase()
	params := &usecase.CreateUserParam{}
	faker.FakeData(params)
	id, _ := fu.CreateUser(context.Background(), params)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users/me", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", id))

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	server := &Server{
		userGetter: fu,
	}

	err := server.GetCurrentUser(ctx)
	assert.NoError(err)

	var response generated.GetUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(err)
	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal(params.FullName, response.FullName)
	assert.Equal(params.PhoneNumber, response.PhoneNumber)
}

func TestGetCurrentUser_UserNotFound(t *testing.T) {
	fu := fake.NewFakeUserUsecase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users/me", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", "1232131"))

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	server := &Server{
		userGetter: fu,
	}

	err := server.GetCurrentUser(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(err)
	assert.Equal(http.StatusNotFound, rec.Code)
	assert.Equal("RecordNotFound", response.Type)
}

func TestGetCurrentUser_ErrOnUsecase(t *testing.T) {
	fu := fake.NewFakeUserUsecase()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users/me", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", "1231232131"))

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	server := &Server{
		userGetter: fu,
	}

	err := server.GetCurrentUser(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(err)
	assert.Equal(http.StatusInternalServerError, rec.Code)
	assert.Equal("InternalServerError", response.Type)
}
