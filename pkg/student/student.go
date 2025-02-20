package student

import (
	"uni_app/models"
	handlers "uni_app/pkg/student/handler"
	repositories "uni_app/pkg/student/repository"
	usecases "uni_app/pkg/student/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *models.Config) {
	roleRepo := repositories.NewStudentRepository(db)
	roleUsecase := usecases.NewStudentUsecase(roleRepo)
	handlers.NewStudentHandler(roleUsecase, e)
}
