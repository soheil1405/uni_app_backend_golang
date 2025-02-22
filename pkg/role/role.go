package role

import (
	handlers "uni_app/pkg/role/handler"
	repositories "uni_app/pkg/role/repository"
	usecases "uni_app/pkg/role/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	roleRepo := repositories.NewRoleRepository(db)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	handlers.NewRoleHandler(roleUsecase, e)
}
