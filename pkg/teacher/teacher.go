package teacher

import (
	handlers "uni_app/pkg/teacher/handler"
	repositories "uni_app/pkg/teacher/repository"
	usecases "uni_app/pkg/teacher/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	teacherRepo := repositories.NewTeacherRepository(db)
	teacherUsecase := usecases.NewTeacherUsecase(teacherRepo, config)
	handlers.NewTeacherHandler(teacherUsecase, e)
}
