package user

import (
	handlers "uni_app/pkg/user/handler"
	repositories "uni_app/pkg/user/repository"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	roleRepo := repositories.NewUserRepository(db)
	roleUsecase := usecases.NewUserUsecase(roleRepo, config)
	handlers.NewUserHandler(roleUsecase, e)
}
