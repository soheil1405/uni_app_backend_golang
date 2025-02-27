package user

import (
	_tokenRepo "uni_app/pkg/token/repository"
	handlers "uni_app/pkg/user/handler"
	repositories "uni_app/pkg/user/repository"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	roleRepo := repositories.NewUserRepository(db)
	tokenRepo := _tokenRepo.NewTokenRepository(db)
	roleUsecase := usecases.NewUserUsecase(roleRepo, tokenRepo, config)
	handlers.NewUserHandler(roleUsecase, e)
}
