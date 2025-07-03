package department

import (
	handlers "uni_app/pkg/department/handler"
	repositories "uni_app/pkg/department/repository"
	usecases "uni_app/pkg/department/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	departmentRepo := repositories.NewDepartmentRepository(db)
	departmentUsecase := usecases.NewDepartmentUsecase(departmentRepo, config)
	handlers.NewDepartmentHandler(departmentUsecase, e)
}
