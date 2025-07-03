package middleware

import (
	"net/http"
	"strings"
	"uni_app/models"
	repositories "uni_app/pkg/auth/repository"
	"uni_app/services/env"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	repo   repositories.AuthRepository
	config *env.AppConfig
}

func NewAuthMiddleware(repo repositories.AuthRepository, config *env.AppConfig) *AuthMiddleware {
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
