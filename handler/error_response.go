package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	customerror "github.com/SawitProRecruitment/UserService/internal/customError"
	"github.com/labstack/echo/v4"
)

func parseError(ctx echo.Context, err error) error {
	switch parsedError := err.(type) {
	case *customerror.ValidationError:
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Type:     "ValidationError",
			Messages: parseItemValidationError(parsedError),
		})
	case *echo.HTTPError:
		return ctx.JSON(parsedError.Code, generated.ErrorResponse{
			Type: "RequestBodyError",
			Messages: []generated.ErrorResponseItem{
				{
					Name:   "request body",
					Reason: parsedError.Internal.Error(),
				},
			},
		})
	default:
		if errors.Is(parsedError, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, generated.ErrorResponse{
				Type: "RecordNotFound",
				Messages: []generated.ErrorResponseItem{
					{
						Name:   "datra",
						Reason: parsedError.Error(),
					},
				},
			})
		}
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Type: "InternalServerError",
			Messages: []generated.ErrorResponseItem{
				{
					Name:   "UnexpectedError",
					Reason: parsedError.Error(),
				},
			},
		})
	}
}

func parseItemValidationError(err *customerror.ValidationError) []generated.ErrorResponseItem {
	var errorItems []generated.ErrorResponseItem

	for _, message := range strings.Split(err.Error(), ";") {
		messages := strings.Split(message, ": ")

		key := messages[0]
		for _, detailError := range strings.Split(messages[1], ",") {
			errorItems = append(errorItems, generated.ErrorResponseItem{
				Name:   key,
				Reason: detailError,
			})
		}
	}
	return errorItems
}
