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

func TestServer_CreateUser_CannotBindBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users",
		strings.NewReader(`{"PhoneNumber": "123456789", "FullName": "John Doe", "Password": "password123",}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert := assert.New(t)

	fu := fake.NewFakeUserUsecase()
	server := &Server{
		uu: fu,
	}

	err := server.CreateUser(ctx)
	assert.NoError(err)

	var response generated.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(err)

	assert.Equal(http.StatusBadRequest, rec.Code)
	assert.Equal("RequestBodyError", response.Type)
}
