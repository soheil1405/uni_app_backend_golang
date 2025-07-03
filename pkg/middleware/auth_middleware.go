package middleware

import (
	"net/http"
	"strings"
	"uni_app/models"
	"uni_app/services/env"
	"uni_app/utils/ctxHelper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	repo   repository.AuthRepository
	config *env.Config
}

func NewAuthMiddleware(repo repository.AuthRepository, config *env.Config) *AuthMiddleware {
	return &AuthMiddleware{
		repo:   repo,
		config: config,
	}
}

func (m *AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization header is required")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			return []byte(m.config.JWTSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string)
			user, err := m.repo.GetByID(userID)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "User not found")
			}

			c.Set("user", user)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}
}

func (m *AuthMiddleware) Role(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*models.User)
			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, "Access denied")
		}
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
