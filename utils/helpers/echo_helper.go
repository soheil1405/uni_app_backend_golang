package helpers

import (
	"errors"
	"net/http"
	"strconv"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

// ParsePaginationParams ...
func ParsePaginationParams(ctx echo.Context) (limit, offset int, err error) {
	limit, _ = strconv.Atoi(ctx.QueryParam("limit"))
	offset, _ = strconv.Atoi(ctx.QueryParam("offset"))

	if limit < 0 {
		err = errors.New("limit must be positive")
	} else if limit == 0 {
		limit = 20
	} else if limit > 500 {
		limit = 500
	}

	return
}

// Reply ...
func Reply(ctx echo.Context, httpStatus int, err error, content map[string]interface{}, meta interface{}) error {
	var template *templates.ResponseTemplate

	switch httpStatus {
	case http.StatusOK:
		template = templates.Ok(content, err, meta)
	case http.StatusCreated:
		template = templates.Created(content, meta)
	case http.StatusBadRequest:
		template = templates.BadRequest(content, err.Error())
	case http.StatusInternalServerError:
		template = templates.InternalServerError(content, err.Error())
	case http.StatusNotFound:
		template = templates.NotFound(content, err.Error())
	case http.StatusUnprocessableEntity:
		template = templates.UnprocessableEntity(content, err.Error())
	case http.StatusMethodNotAllowed:
		template = templates.MethodNotAllowed(content, err.Error())
	case http.StatusUnauthorized:
		template = templates.Unauthorized(content, err.Error())
	case http.StatusForbidden:
		template = templates.Forbidden(content, err.Error())
	case http.StatusGatewayTimeout:
		template = templates.GatewayTimeOut(content, err.Error())
	case http.StatusLocked:
		template = templates.Locked(content, err.Error())
	case http.StatusNotAcceptable:
		template = templates.NotAcceptable(content, err.Error())
	default:
		template = templates.InternalServerError(content, errors.New("invalid reply request"))
	}

	return ctx.JSON(httpStatus, template)
}
