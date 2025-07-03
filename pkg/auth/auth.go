package auth

import (
	handlers "uni_app/pkg/auth/handler"
	repositories "uni_app/pkg/auth/repository"
	usecases "uni_app/pkg/auth/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	authRepo := repositories.NewAuthRepository(db)
	authUsecase := usecases.NewAuthUsecase(authRepo, config)
	handlers.NewAuthHandler(authUsecase, e)
}
