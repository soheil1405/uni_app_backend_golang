package ctxHelper

import (
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

const (
	UserContextKey = "user"
)

// SetUserInContext sets the user in the echo context
func SetUserInContext(c echo.Context, user *models.User) {
	c.Set(UserContextKey, user)
}

// GetUserFromContext gets the user from the echo context
func GetUserFromContext(c echo.Context) *models.User {
	user, ok := c.Get(UserContextKey).(*models.User)
	if !ok {
		return nil
	}
	return user
}

// GetIDFromContext gets the ID from the echo context
func GetIDFromContext(c echo.Context) (uint, error) {
	id := c.Param("id")
	if id == "" {
		return 0, echo.NewHTTPError(400, "id is required")
	}
	return uint(helpers.ParseUint(id)), nil
}
