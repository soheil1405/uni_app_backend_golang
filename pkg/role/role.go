package role

import (
	"uni_app/models"
	handlers "uni_app/pkg/role/handler"
	repositories "uni_app/pkg/role/repository"
	usecases "uni_app/pkg/role/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *models.Config) {
	roleRepo := repositories.NewRoleRepository(db)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	handlers.NewRoleHandler(roleUsecase, e)
}
