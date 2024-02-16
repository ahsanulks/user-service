package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"userservice/generated"
	"userservice/internal/user/param/request"
	"userservice/test/fake"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_UpdateCurrentUser_CannotBindBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/users/me",
		strings.NewReader(`{"PhoneNumber": "123456789", "FullName": "John Doe",}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.UpdateCurrentUserProfile(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusBadRequest, rec.Code)
	assert.Equal("RequestBodyError", response.Type)
}

func TestServer_UpdateCurrentUser_UnexpectedError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/users/me",
		strings.NewReader(`{"PhoneNumber": "11111111","FullName": "John Doe"}`),
	)
	req = req.WithContext(context.WithValue(req.Context(), "userID", "1232131"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.UpdateCurrentUserProfile(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusInternalServerError, rec.Code)
	assert.Equal("InternalServerError", response.Type)
}

func TestServer_UpdateCurrentUser_DuplicateError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/users/me",
		strings.NewReader(`{"PhoneNumber": "11111111","FullName": "John Doe"}`),
	)
	req = req.WithContext(context.WithValue(req.Context(), "userID", "3333"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		userUsecase: fu,
	}

	err := server.UpdateCurrentUserProfile(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusConflict, rec.Code)
	assert.Equal("DuplicateResource", response.Type)
}

func TestServer_UpdateCurrentUser_Success(t *testing.T) {
	fu := fake.NewFakeUserUsecase()
	params := &request.CreateUser{}
	faker.FakeData(params)
	id, _ := fu.CreateUser(context.Background(), params)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/users/me",
		strings.NewReader(`{"PhoneNumber": "11111111","FullName": "John Doe"}`),
	)
	req = req.WithContext(context.WithValue(req.Context(), "userID", id))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	server := &Server{
		userUsecase: fu,
	}

	err := server.UpdateCurrentUserProfile(ctx)
	assert.NoError(err)

	var response generated.UpdateUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(http.StatusOK, rec.Code)
	assert.NoError(err)
	assert.Equal(params.FullName, response.FullName)
	assert.Equal(params.PhoneNumber, response.PhoneNumber)
}
