package student

import (
	handlers "uni_app/pkg/student/handler"
	repositories "uni_app/pkg/student/repository"
	usecases "uni_app/pkg/student/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	studentRepo := repositories.NewStudentRepository(db)
	studentUsecase := usecases.NewStudentUsecase(studentRepo, config)
	handlers.NewStudentHandler(studentUsecase, e)
}
