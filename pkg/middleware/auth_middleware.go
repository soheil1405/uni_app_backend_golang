package middleware

import (
	"net/http"
	"strings"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	userUsecase usecases.UserUsecase
}

func NewAuthMiddleware(userUsecase usecases.UserUsecase) *AuthMiddleware {
	return &AuthMiddleware{userUsecase}
}

func (m *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "authorization header is required"})
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header format"})
		}

		// Get the token
		token := parts[1]

		// Validate the token and get the user
		user, err := m.userUsecase.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		// Set the user in the context
		ctxHelper.SetUserInContext(c, user)

		return next(c)
	}
}

func (m *AuthMiddleware) RequireRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := ctxHelper.GetUserFromContext(c)
			if user == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not found in context"})
			}

			// Check if user has any of the required roles
			for _, role := range roles {
				for _, userRole := range user.UserRoles {
					if userRole.Role.Name == role {
						return next(c)
					}
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"error": "insufficient permissions"})
		}
	}
}

func (m *AuthMiddleware) RequireStudent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := ctxHelper.GetUserFromContext(c)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not found in context"})
		}

		for _, userRole := range user.UserRoles {
			if userRole.Role.Name == "student" {
				return next(c)
			}
		}

		return c.JSON(http.StatusForbidden, map[string]string{"error": "only students can access this resource"})
	}
}

func (m *AuthMiddleware) RequireTeacher(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := ctxHelper.GetUserFromContext(c)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not found in context"})
		}

		for _, userRole := range user.UserRoles {
			if userRole.Role.Name == "teacher" {
				return next(c)
			}
		}

		return c.JSON(http.StatusForbidden, map[string]string{"error": "only teachers can access this resource"})
	}
}

func (m *AuthMiddleware) RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := ctxHelper.GetUserFromContext(c)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not found in context"})
		}

		for _, userRole := range user.UserRoles {
			if userRole.Role.Name == "admin" {
				return next(c)
			}
		}

		return c.JSON(http.StatusForbidden, map[string]string{"error": "only admins can access this resource"})
	}
}
