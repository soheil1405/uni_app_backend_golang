package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"uni_app/models"
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

// ParsePaginationParams ...
func BindFetchRequestFromCtx(ctx echo.Context) (*models.FetchRequest, MyError) {
	var (
		MyErr   MyError
		request = &models.FetchRequest{}
	)
	MyErr.Default()

	if err := ctx.Bind(request); err != nil {
		MyErr.SetError(err, http.StatusBadRequest)
		return nil, MyErr
	}

	filters := make(map[string]interface{})
	for key, values := range ctx.QueryParams() {
		if strings.HasPrefix(key, "filters[") && strings.HasSuffix(key, "]") {
			filterKey := key[8 : len(key)-1] // Extract key inside brackets
			if len(values) == 1 {
				filters[filterKey] = values[0]
			} else {
				filters[filterKey] = values
			}
		}
	}
	request.Filters = filters

	if request.Limit <= 0 {
		request.Limit = 20
	} else if request.Limit >= 100 {
		request.Limit = 100
	}

	if request.Offset < 0 {
		request.Offset = 0
	}

	if request.Page > 0 && request.Offset == 0 {
		request.Offset = (request.Page - 1) * request.Limit
	}

	validSorts := []string{}
	for _, sort := range request.Sorts {
		if strings.HasPrefix(sort, "-") {
			validSorts = append(validSorts, fmt.Sprintf("%s DESC", strings.TrimPrefix(sort, "-")))
		} else {
			validSorts = append(validSorts, fmt.Sprintf("%s ASC", sort))
		}
	}
	request.Sorts = validSorts

	if request.ID.IsValid() {
		request.IDs = append(request.IDs, request.ID)
	}

	return request, MyErr
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
