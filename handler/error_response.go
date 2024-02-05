package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	customerror "github.com/SawitProRecruitment/UserService/internal/customError"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func parseValidationError(ctx echo.Context, err *customerror.ValidationError) error {
	return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
		Type:     "ValidationError",
		Messages: parseItemValidationError(err),
	})
}

func parseHTTPError(ctx echo.Context, err *echo.HTTPError) error {
	return ctx.JSON(err.Code, generated.ErrorResponse{
		Type: "RequestBodyError",
		Messages: []generated.ErrorResponseItem{
			{
				Name:   "request body",
				Reason: err.Internal.Error(),
			},
		},
	})
}

func parsePQError(ctx echo.Context, err *pq.Error) error {
	if err.Code == "23505" {
		return ctx.JSON(http.StatusConflict, generated.ErrorResponse{
			Type: "DuplicateResource",
			Messages: []generated.ErrorResponseItem{
				{
					Name:   "resource",
					Reason: "duplicate violation on unique constraint",
				},
			},
		})
	}

	return parseDefaultError(ctx, err)
}

func parseNoRowsError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusNotFound, generated.ErrorResponse{
		Type: "RecordNotFound",
		Messages: []generated.ErrorResponseItem{
			{
				Name:   "data",
				Reason: err.Error(),
			},
		},
	})
}

func parseDefaultError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{
		Type: "InternalServerError",
		Messages: []generated.ErrorResponseItem{
			{
				Name:   "UnexpectedError",
				Reason: err.Error(),
			},
		},
	})
}

func parseError(ctx echo.Context, err error) error {
	switch parsedError := err.(type) {
	case *customerror.ValidationError:
		return parseValidationError(ctx, parsedError)
	case *echo.HTTPError:
		return parseHTTPError(ctx, parsedError)
	case *pq.Error:
		return parsePQError(ctx, parsedError)
	default:
		if errors.Is(parsedError, sql.ErrNoRows) {
			return parseNoRowsError(ctx, parsedError)
		}
		return parseDefaultError(ctx, parsedError)
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
